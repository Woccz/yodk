package nolol

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dbaumgarten/yodk/pkg/nolol/nast"
	"github.com/dbaumgarten/yodk/pkg/optimizers"
	"github.com/dbaumgarten/yodk/pkg/parser"
	"github.com/dbaumgarten/yodk/pkg/parser/ast"
)

// Converter can convert a nolol-ast to a yolol-ast
type Converter struct {
	files      FileSystem
	lineLabels map[string]int
	// the names of definitions are case-insensitive. Keys are converted to lowercase before using them
	// all lookups MUST also use lowercased keys
	definitions      map[string]*nast.Definition
	usesTimeTracking bool
	iflabelcounter   int
	waitlabelcounter int
	loopcounter      int
	// keeps track of the current loop we are in while converting
	// the last element in the list is the current innermost loop
	loopLevel           []int
	sexpOptimizer       *optimizers.StaticExpressionOptimizer
	boolexpOptimizer    *optimizers.ExpressionInversionOptimizer
	varnameOptimizer    *optimizers.VariableNameOptimizer
	includecount        int
	macros              map[string]*nast.MacroDefinition
	macroLevel          []string
	macroInsertionCount int
	debug               bool
	// Spaceless uses spaceless printer-style for yolol
	Spaceless bool
}

// NewConverter creates a new converter
func NewConverter() *Converter {
	return &Converter{
		lineLabels:       make(map[string]int),
		definitions:      make(map[string]*nast.Definition),
		macros:           make(map[string]*nast.MacroDefinition),
		macroLevel:       make([]string, 0),
		sexpOptimizer:    optimizers.NewStaticExpressionOptimizer(),
		boolexpOptimizer: &optimizers.ExpressionInversionOptimizer{},
		varnameOptimizer: optimizers.NewVariableNameOptimizer(),
		loopLevel:        make([]int, 0),
	}
}

// GetVariableTranslations returns a table that can be used to find the original names
// of the variables whos names where shortened during conversion
func (c *Converter) GetVariableTranslations() map[string]string {
	return c.varnameOptimizer.GetReversalTable()
}

// ConvertFile is a shortcut that loads a file from the file-system, parses it and directly convertes it.
// mainfile is the path to the file on the disk.
// All included are loaded relative to the mainfile.
func (c *Converter) ConvertFile(mainfile string) (*ast.Program, error) {
	files := DiskFileSystem{
		Dir: filepath.Dir(mainfile),
	}
	return c.ConvertFileEx(filepath.Base(mainfile), files)
}

// ConvertFileEx acts like ConvertFile, but allows the passing of a custom filesystem from which the source files
// are retrieved. This way, files that are not stored on disk can be converted
func (c *Converter) ConvertFileEx(mainfile string, files FileSystem) (*ast.Program, error) {
	file, err := files.Get(mainfile)
	if err != nil {
		return nil, err
	}
	p := NewParser()
	p.Debug(c.debug)
	parsed, err := p.Parse(file)
	if err != nil {
		return nil, err
	}
	return c.Convert(parsed, files)
}

// Debug enables/disables debug logging
func (c *Converter) Debug(b bool) {
	c.debug = b
}

// Convert converts a nolol-program to a yolol-program
// files is an object to access files that are referenced in prog's include directives
func (c *Converter) Convert(prog *nast.Program, files FileSystem) (*ast.Program, error) {
	c.files = files

	c.usesTimeTracking = usesTimeTracking(prog)
	// reserve a name for use in time-tracking
	c.varnameOptimizer.OptimizeVarName(reservedTimeVariable)

	// find all user-defined line-labels
	err := c.findLineLabels(prog, false)
	if err != nil {
		return nil, err
	}

	err = c.convertNodes(prog)
	if err != nil {
		return nil, err
	}

	err = c.addFinalGoto(prog)
	if err != nil {
		return nil, err
	}

	err = c.resolveGotoChains(prog)
	if err != nil {
		return nil, err
	}

	err = c.removeUnusedLabels(prog)
	if err != nil {
		return nil, err
	}

	// merge the statemens of the program as good as possible
	merged, err := c.mergeNololElements(prog.Elements)
	if err != nil {
		return nil, err
	}
	prog.Elements = merged

	err = c.removeDuplicateGotos(prog)
	if err != nil {
		return nil, err
	}

	// find all line-labels (again). This time they have the correct lines.
	err = c.findLineLabels(prog, true)
	if err != nil {
		return nil, err
	}

	// resolve line-labels
	err = c.replaceLineLabels(prog)
	if err != nil {
		return nil, err
	}

	// convertLineFuncCalls might have introduced un-optimized expression
	// re-run the static-expression optimizer
	err = c.sexpOptimizer.Optimize(prog)
	if err != nil {
		return nil, err
	}

	if c.usesTimeTracking {
		c.insertLineCounter(prog)
	}

	// at this point the program consists entirely of statement-lines which contain pure yolol-code
	out := &ast.Program{
		Lines: make([]*ast.Line, len(prog.Elements)),
	}

	for i, element := range prog.Elements {
		line := element.(*nast.StatementLine)
		out.Lines[i] = &ast.Line{
			Position:   line.Position,
			Statements: line.Statements,
		}
	}

	c.removeFinalGotoIfNeeded(out)

	if len(out.Lines) > 20 {
		return out, &parser.Error{
			Message: "Program is too large to be compiled into 20 lines of yolol.",
			StartPosition: ast.Position{
				Line:    1,
				Coloumn: 1,
			},
			EndPosition: ast.Position{
				Line:    30,
				Coloumn: 70,
			},
		}
	}

	return out, nil
}

func (c *Converter) maxLineLength() int {
	if !c.usesTimeTracking {
		return 70
	}
	return 70 - 4
}

func (c *Converter) convertNodes(node ast.Node) error {
	f := func(node ast.Node, visitType int) error {
		switch n := node.(type) {
		case *ast.Assignment:
			return c.convertAssignment(n, visitType)

		case *nast.Definition:
			return c.convertDefinition(n, visitType)

		case *nast.MacroDefinition:
			return c.convertMacroDef(n, visitType)

		case *nast.MacroInsetion:
			return c.convertMacroInsertion(n, visitType)

		case *nast.IncludeDirective:
			return c.convertInclude(n)

		case *nast.WaitDirective:
			return c.convertWait(n, visitType)

		case *nast.FuncCall:
			return c.convertFuncCall(n, visitType)

		case *ast.Dereference:
			return c.convertDereference(n)

		case *nast.MultilineIf:
			return c.convertIf(n, visitType)

		case *nast.WhileLoop:
			return c.convertWhileLoop(n, visitType)

		case *nast.BreakStatement:
			return c.convertBreakStatement(n)

		case *nast.ContinueStatement:
			return c.convertContinueStatement(n)

		case *ast.UnaryOperation:
			if visitType == ast.PostVisit {
				return ast.NewNodeReplacementSkip(c.optimizeExpression(n))
			}
		case *ast.BinaryOperation:
			if visitType == ast.PostVisit {
				return ast.NewNodeReplacementSkip(c.optimizeExpression(n))
			}
		case *nast.Trigger:
			if n.Kind == "macroleft" {
				c.macroLevel = c.macroLevel[:len(c.macroLevel)-1]
				return ast.NewNodeReplacement()
			}
		}

		return nil
	}
	return node.Accept(ast.VisitorFunc(f))
}

func (c *Converter) optimizeExpression(exp ast.Expression) ast.Node {
	repl := c.boolexpOptimizer.OptimizeExpression(exp)
	if repl != nil {
		exp = repl
	}
	repl = c.sexpOptimizer.OptimizeExpressionNonRecursive(exp)
	if repl != nil {
		exp = repl
	}
	return exp
}

// mergeNololNestableElements is a type-wrapper for mergeStatementElements
func (c *Converter) mergeNololNestableElements(lines []nast.NestableElement) ([]nast.NestableElement, error) {
	inp := make([]*nast.StatementLine, len(lines))
	for i, elem := range lines {
		line, isline := elem.(*nast.StatementLine)
		if !isline {
			return nil, parser.Error{
				Message: fmt.Sprintf("Err: Found unconverted nolol-element: %T", elem),
			}
		}
		inp[i] = line
	}
	interm, err := c.mergeStatementElements(inp)
	if err != nil {
		return nil, err
	}
	outp := make([]nast.NestableElement, len(interm))
	for i, elem := range interm {
		outp[i] = elem
	}
	return outp, nil
}

// mergeNololElements is a type-wrapper for mergeStatementElements
func (c *Converter) mergeNololElements(lines []nast.Element) ([]nast.Element, error) {
	inp := make([]*nast.StatementLine, len(lines))
	for i, elem := range lines {
		line, isline := elem.(*nast.StatementLine)
		if !isline {
			return nil, parser.Error{
				Message: fmt.Sprintf("Err: Found unconverted nolol-element: %T", elem),
			}
		}
		inp[i] = line
	}
	interm, err := c.mergeStatementElements(inp)
	if err != nil {
		return nil, err
	}
	outp := make([]nast.Element, len(interm))
	for i, elem := range interm {
		outp[i] = elem
	}
	return outp, nil
}

// mergeStatementElements merges consectuive statementlines into as few lines as possible
func (c *Converter) mergeStatementElements(lines []*nast.StatementLine) ([]*nast.StatementLine, error) {
	maxlen := c.maxLineLength()
	newElements := make([]*nast.StatementLine, 0, len(lines))
	i := 0
	for i < len(lines) {
		current := &nast.StatementLine{
			Line: ast.Line{
				Statements: []ast.Statement{},
				Position:   lines[i].Position,
			},
			Label:  lines[i].Label,
			HasEOL: lines[i].HasEOL,
		}
		current.Statements = append(current.Statements, lines[i].Statements...)
		newElements = append(newElements, current)

		if current.HasEOL {
			// no lines MUST be appended to a line having EOL
			i++
			continue
		}

		for i+1 < len(lines) {
			currlen := c.getLengthOfLine(&current.Line)

			if currlen > maxlen {
				return newElements, &parser.Error{
					Message:       "The line is too long (>70 characters) to be converted to yolol, even after optimization.",
					StartPosition: current.Start(),
					EndPosition:   current.End(),
				}
			}

			nextline := lines[i+1]

			if nextline.Label == "" && !nextline.HasBOL {
				prev := current.Statements
				current.Statements = make([]ast.Statement, 0, len(current.Statements)+len(nextline.Statements))
				current.Statements = append(current.Statements, prev...)
				current.Statements = append(current.Statements, nextline.Statements...)

				newlen := c.getLengthOfLine(&current.Line)
				if newlen > maxlen {
					// the newly created line is longer then allowed. roll back.
					current.Statements = prev
					break
				}

				i++
				if nextline.HasEOL {
					break
				}
			} else {
				break
			}
		}
		i++
	}
	return newElements, nil
}

//getLengthOfLine returns the amount of characters needed to represent the given line as yolol-code
func (c *Converter) getLengthOfLine(line ast.Node) int {
	ygen := parser.Printer{}
	if c.Spaceless {
		ygen.Mode = parser.PrintermodeSpaceless
	} else {
		ygen.Mode = parser.PrintermodeCompact
	}

	silenceGotoExpression := false
	ygen.PrinterExtensionFunc = func(node ast.Node, visitType int, p *parser.Printer) (bool, error) {
		if gotostmt, is := node.(*ast.GoToStatement); is {
			if c.getGotoDestinationLabel(gotostmt) != "" {
				if visitType == ast.PreVisit {
					silenceGotoExpression = true
					p.Write("gotoXX")
				}
				if visitType == ast.PostVisit {
					silenceGotoExpression = false
				}
				return true, nil
			}
		}
		if silenceGotoExpression {
			if _, is := node.(ast.Expression); is {
				// The current expression is inside a goto.
				// DO NOT PRINT IT
				return true, nil
			}
		}
		return false, nil
	}
	generated, err := ygen.Print(line)
	if err != nil {
		panic(err)
	}

	linelen := len(generated)
	if strings.HasSuffix(generated, "\n") {
		linelen--
	}

	return linelen
}

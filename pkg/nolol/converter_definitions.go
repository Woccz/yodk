package nolol

import (
	"strings"

	"github.com/dbaumgarten/yodk/pkg/nolol/nast"
	"github.com/dbaumgarten/yodk/pkg/parser"
	"github.com/dbaumgarten/yodk/pkg/parser/ast"
)

// getDefinition is a case-insensitive getter for c.definitions
func (c *Converter) getDefinition(name string) (*nast.Definition, bool) {
	name = strings.ToLower(name)
	val, exists := c.definitions[name]
	return val, exists
}

// setDefinition is a case-insensitive setter for c.definitions
func (c *Converter) setDefinition(name string, val *nast.Definition) {
	name = strings.ToLower(name)
	c.definitions[name] = val
}

// convertDefinitions converts a definition to yolol by discarding it, but saving the defined value
func (c *Converter) convertDefinition(decl *nast.Definition, visitType int) error {
	if visitType == ast.PreVisit {
		c.setDefinition(decl.Name, decl)
		return ast.NewNodeReplacement()
	}
	return nil
}

// convertDefinitionFunction converts a definition that contains placeholders (=behaves like a function)
func (c *Converter) convertDefinitionFunction(fc *nast.FuncCall) error {
	def, exists := c.getDefinition(fc.Function)
	if !exists {
		return nil
	}
	if len(def.Placeholders) == 0 {
		return nil
	}

	if len(def.Placeholders) != len(fc.Arguments) {
		return nil
	}

	// gather replacements
	replacements := make(map[string]ast.Expression)
	for i := range fc.Arguments {
		lvarname := strings.ToLower(def.Placeholders[i])
		replacements[lvarname] = fc.Arguments[i]
	}

	copy := nast.CopyAst(def.Value)
	err := c.replacePlaceholders(copy, replacements, nil, false)
	if err != nil {
		return err
	}

	return ast.NewNodeReplacement(copy)
}

// convertDefinitionAssignment replaces assignments with definitions
func (c *Converter) convertDefinitionAssignment(ass *ast.Assignment, visitType int) error {
	if visitType == ast.PostVisit {
		if replacement, exists := c.getDefinition(ass.Variable); exists {
			if replacementVariable, isvar := replacement.Value.(*ast.Dereference); isvar && replacementVariable.Operator == "" {
				ass.Variable = replacementVariable.Variable
			} else {
				return &parser.Error{
					Message:       "Can not assign to a definition that is an expression (need a single variable name)",
					StartPosition: ass.Start(),
					EndPosition:   ass.End(),
				}
			}
		}
	}
	return nil
}

// convertDereference replaces mentionings of definitions with the value of the definition
func (c *Converter) convertDefinitionDereference(deref *ast.Dereference) error {
	if definition, exists := c.getDefinition(deref.Variable); exists {
		// dereference of definition
		replacement := nast.CopyAst(definition.Value)
		if replacementVariable, isvar := replacement.(*ast.Dereference); isvar {
			if deref.Operator != "" && replacementVariable.Operator != "" {
				return &parser.Error{
					Message:       "You can not use pre/post-operators on defitions that use the operator themselves",
					StartPosition: deref.Start(),
					EndPosition:   deref.End(),
				}
			}
			if deref.Operator != "" {
				replacementVariable.Operator = deref.Operator
				replacementVariable.PrePost = deref.PrePost
			}
		} else if deref.Operator != "" {
			return &parser.Error{
				Message:       "Can not use pre/port-operators on expressions",
				StartPosition: deref.Start(),
				EndPosition:   deref.End(),
			}
		}
		return ast.NewNodeReplacement(replacement)
	}
	return nil
}

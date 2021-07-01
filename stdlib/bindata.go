// Code generated for package stdlib by go-bindata DO NOT EDIT. (@generated)
// sources:
// src/logic.nolol
// src/math_advanced.nolol
// src/math_basic.nolol
// src/math_professional.nolol
// src/string.nolol
package stdlib

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _logicNolol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x4f\x6f\xd3\x4e\x10\xbd\xfb\x53\xbc\x5f\x54\xa9\xc9\x8f\x04\x37\x1c\x29\xe1\xc0\xad\x57\x04\xe7\x6a\xd7\x1e\xdb\xa3\xda\x33\xd6\xec\x2c\x29\xdf\x1e\xad\x43\xd3\xa4\xe5\x64\xc9\xf3\xf4\xfe\x6e\x5d\xe3\xc7\xc0\x09\x1d\x8f\x84\x46\xc5\x03\x4b\x42\x0c\x89\x1b\xb4\xd4\xb1\xb0\xb3\x4a\x42\x90\x16\x53\x68\x4c\x13\x3a\x35\x8c\xda\x73\xb3\xd3\x99\x2c\x2c\xf7\xaa\xae\xf1\x30\xcd\x6a\x8e\x9c\x58\x7a\xdc\x82\xa5\x19\x73\x4b\x58\x25\x6f\xeb\x05\xbf\xc2\x6d\x55\x80\xdf\xc9\xb3\x15\x4e\x70\x57\x34\xdb\x45\x03\x9c\xe0\x96\x69\x0b\xf5\x81\xec\xc8\x89\x10\x0b\xfc\x8c\xd8\x22\x2c\x3e\x22\xa6\x9c\x1c\x91\x20\x79\x8a\x64\xe9\xe3\x05\xcb\xcb\xe9\x0e\x6a\xd8\x57\x8b\xe7\x93\xdd\x47\x27\x93\x60\xbf\xd7\x97\x84\x5b\xc4\x0d\xe8\x79\xb6\x0a\x00\x22\x3e\x60\x1d\x76\x71\xf3\xff\x19\x53\x91\xb4\x8b\xeb\x87\x37\x5e\xef\xb6\x98\x4d\xdb\xdc\x50\x49\x62\x59\x9c\x27\xda\x91\x99\x1a\x7c\x08\x8e\x23\x8f\x23\xd2\x13\xcf\xf0\x81\x60\x34\x05\x96\x52\xcd\xc8\x42\x85\xf0\x67\x0a\x3d\x7d\xc6\xea\x64\xae\x74\xcf\x92\xe9\xb1\x9c\xd7\xbf\x82\x6d\xee\xd1\xea\x61\x7f\x8f\xe4\xb9\xeb\x0e\x9f\x70\xb3\xaa\x96\xb5\x08\x37\x45\x9f\x97\xba\x83\x38\x5c\x21\xea\x7f\xa5\x54\x31\xe5\x66\xb8\x0a\x7e\xcd\x7d\x4e\xb1\xf9\xc2\xbd\xa8\xd1\xd7\x93\xa5\x52\xc0\xe9\x47\x7d\x78\x1f\xff\x65\xb4\x7d\x19\x2d\x94\x72\xe3\xeb\x62\x31\xfb\x62\x21\xaa\x0f\x05\xfc\x76\xa7\x7f\x8c\xf1\xac\xb6\x0e\xdb\xcb\xf2\x03\xfe\x3b\x20\x9e\x05\xbf\x8d\xda\x3c\x25\x84\x84\x51\xa5\x2f\xdf\x77\x4f\xe5\x8a\xf0\x18\xd8\x2f\xb2\xbd\x66\x1a\xa8\x44\xbc\x7a\x6a\x3e\x90\xa0\x57\xd7\xe5\x88\x22\x49\xd2\xfe\x09\x00\x00\xff\xff\x58\x96\xdd\x23\x0c\x03\x00\x00")

func logicNololBytes() ([]byte, error) {
	return bindataRead(
		_logicNolol,
		"logic.nolol",
	)
}

func logicNolol() (*asset, error) {
	bytes, err := logicNololBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logic.nolol", size: 780, mode: os.FileMode(438), modTime: time.Unix(1623961616, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _math_advancedNolol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xd0\xc1\x6a\xc3\x30\x0c\x06\xe0\xf3\xfc\x14\x3f\x85\xd2\x14\xb6\x85\x3d\xc6\xae\x63\xf7\xe2\xc4\x4a\x23\xb0\xa5\x60\x29\x9d\x1f\x7f\x34\x2b\xac\x90\xa3\x84\xf4\xfd\xf0\xf7\x3d\xbe\x67\x36\x4c\x9c\x09\xa3\x8a\x47\x16\xc3\x10\x8d\x47\x24\x9a\x58\xd8\x59\xc5\x10\x25\xa1\xc4\xb1\xaa\x61\xd2\x8a\x12\x7d\x0e\x7d\x8f\xcf\xb2\x68\x75\xac\xc6\x72\xc5\x09\x2c\x63\x5e\x13\xe1\x60\x9e\xfa\xfb\xcd\x01\xa7\x10\x76\xdb\xcb\xe6\x1f\xc2\x5d\xf8\x22\x5f\xab\x18\x7c\x26\x08\x35\x47\xd6\x1f\xaa\x60\x71\xba\x52\x85\x2b\x5a\xd8\x82\xb7\xcc\xcb\x94\x55\x6b\xd7\xce\xa0\xb6\xd4\xf0\xd2\xde\xda\xf1\x23\x90\xa4\x9d\x15\x07\xd3\xbc\x3a\xe1\x16\xf3\x4a\xd0\x09\xed\xfd\x19\x8a\x83\xfd\x33\x7f\xc3\xce\x69\xc7\x82\x8e\x6e\x24\x50\x79\x74\x32\xce\xbc\xd8\xf9\x19\x2a\x9a\xba\xf6\x5a\x1e\x14\x80\xed\x6d\xb3\x7e\x03\x00\x00\xff\xff\xe4\x15\xcd\xd9\x5d\x01\x00\x00")

func math_advancedNololBytes() ([]byte, error) {
	return bindataRead(
		_math_advancedNolol,
		"math_advanced.nolol",
	)
}

func math_advancedNolol() (*asset, error) {
	bytes, err := math_advancedNololBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math_advanced.nolol", size: 349, mode: os.FileMode(438), modTime: time.Unix(1625166930, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _math_basicNolol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x41\x6f\xa3\x30\x10\x46\xcf\xcb\xaf\xf8\x14\x69\x15\x88\x92\x60\x76\x57\xda\x1e\x42\xee\xb9\x56\xbd\x57\x0e\x0c\x61\x24\x6c\x23\x7b\x48\xfc\xf3\x2b\x08\x4a\x68\x7b\xe8\x05\x64\x79\xde\x7b\xb6\x9c\xe7\x78\x6b\x39\xa0\xe1\x8e\x50\x39\x2b\x9a\x6d\xc0\x59\x07\xae\x50\x53\xc3\x96\x85\x9d\x0d\xd0\xb6\x86\xd1\x95\x77\x01\x8d\xf3\x30\x5a\xda\x24\xcf\x71\x32\xbd\xf3\x82\x21\xb0\xbd\x60\x0d\xb6\x55\x37\xd4\x84\x55\x90\x3a\x1f\x67\x56\x58\x27\xc9\xd4\xa0\x89\x21\xa3\x85\x2b\xdd\x8d\xa9\x20\xda\x0a\x7a\x4e\xa6\xce\x7d\xff\xbd\xe7\xf2\xef\xbe\xf8\x57\xfc\x40\xd1\x27\x88\xca\x3f\xfb\xff\xc5\xcb\xc4\xbc\x92\x0c\xde\x06\x48\x4b\xd0\xe7\xe0\xba\x41\x08\x57\xdd\x0d\x04\xd7\x20\xee\x93\xe9\x16\x77\x4e\x9f\x43\x1a\x33\x50\xec\x7d\xf2\x2b\x8d\xc7\x52\x65\x9b\xb8\x4b\xe3\x61\xfc\x27\x64\xeb\xc9\x78\x6a\x10\xc1\x01\x47\xb5\x85\x9f\xed\x05\xb8\xc1\x41\xcd\x6b\xec\x8a\x2d\x9c\xb4\xe4\x6f\x1c\xe8\x31\xa4\x96\xad\xc0\x17\xfb\x35\x36\xa6\x4a\x95\x3d\x4a\xcb\xb3\x5b\x8a\x82\xce\xdd\xc8\x83\xad\xd0\x85\x3c\xc4\x21\x2e\x95\x4d\xe7\x9c\x7f\x3a\x63\x5e\x28\xa5\x36\xe3\xe7\x9b\x31\xfe\x36\x48\xe9\x4a\x16\xce\xce\xaf\x5b\xb5\xdc\x87\x6c\xe9\x33\xae\x4e\xe3\xd6\xcc\x3e\x00\x88\xbb\x98\x9b\xa7\x76\x63\x46\xf1\x47\x00\x00\x00\xff\xff\xcf\x1d\xb0\x84\x33\x02\x00\x00")

func math_basicNololBytes() ([]byte, error) {
	return bindataRead(
		_math_basicNolol,
		"math_basic.nolol",
	)
}

func math_basicNolol() (*asset, error) {
	bytes, err := math_basicNololBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math_basic.nolol", size: 563, mode: os.FileMode(438), modTime: time.Unix(1625166940, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _math_professionalNolol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xca\x31\x0e\xc2\x30\x0c\x05\xd0\x3d\xa7\xf8\xea\xd2\xd1\xe7\x60\x67\x47\xc6\x4e\xa8\xa5\xc6\x46\xb1\xcb\xf9\x11\x33\xeb\xd3\x23\xc2\xfd\xb0\xc4\xb0\xb3\x43\xc2\x8b\xcd\x13\x4f\x4e\x13\x68\x1f\xe6\x56\x16\x9e\x60\x57\x4c\x96\x15\x89\x11\x0b\x93\xeb\x68\x44\xb8\xcd\x77\xac\xc2\x95\xe6\x2f\xec\x30\x97\xf3\xd2\x8e\x2d\x4b\xe9\x77\x36\xec\xad\xfd\xe9\x83\xf5\xc3\x2e\x5d\xb7\xf6\x0d\x00\x00\xff\xff\xc8\xe5\x98\x34\x80\x00\x00\x00")

func math_professionalNololBytes() ([]byte, error) {
	return bindataRead(
		_math_professionalNolol,
		"math_professional.nolol",
	)
}

func math_professionalNolol() (*asset, error) {
	bytes, err := math_professionalNololBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math_professional.nolol", size: 128, mode: os.FileMode(438), modTime: time.Unix(1625163935, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stringNolol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x31\x8f\xd4\x30\x10\x85\xeb\xcb\xaf\x78\x97\xe6\x58\x2d\x21\xa2\x3e\x16\x89\x92\x16\xd1\x23\xe3\x4c\x92\x91\x92\xb1\x35\x33\x39\xe5\xe7\x23\x27\x61\x51\xd0\x52\x50\x3e\xfb\x79\xde\xf7\x3c\x6d\x8b\xef\x23\x1b\x7a\x9e\x08\x31\x89\x07\x16\xc3\xcf\x60\x1c\xd1\x51\xcf\xc2\xce\x49\x0c\x41\x3a\xcc\x21\x6a\x32\xf4\x49\x61\xae\x2c\x43\x33\x07\xe1\xbc\x4c\xa1\x78\xaa\xb6\xc5\xd7\x39\x27\x75\x2c\xc6\x32\xe0\x05\x2c\x71\x5a\x3a\x42\x6d\xde\xb5\xfb\x93\x1a\x2f\x55\x71\x7e\xa3\x39\xbd\x91\xc1\x47\xc2\x14\xcc\x11\xc7\xa0\x21\x3a\x29\x7a\x4d\x73\x09\xd8\x32\xf3\x14\x22\x19\xd8\xc1\xe2\x09\x69\xf1\x6a\xc3\x38\x08\x7e\xe4\x94\xdf\x99\xeb\xfb\x72\x73\xf9\xc4\x83\x24\xa5\xcf\x98\x58\xa8\x7a\xda\xd5\xcd\x5c\x5f\xcb\xf5\x6d\xd7\x8d\xb9\x36\x0d\x2a\x92\xee\x20\xf1\x45\xc5\xf0\x11\xdc\x6f\xb1\xf7\x5f\x58\xcf\x51\xbf\xcf\xf7\xbc\xf5\x02\x5a\xb3\x56\x4f\x45\x36\xeb\xe5\xb9\xe4\xdc\x87\x7e\xe9\xba\xa3\x1b\xc9\x30\x3a\xd2\x3e\xfa\x68\xd0\xb6\x9b\x62\x83\x91\x97\xc3\xba\x06\xcb\xe6\xcf\x9a\x22\x99\x9d\x93\x27\x92\x3f\x25\x8f\x72\x23\x95\xa2\x3b\xf2\xf3\xad\xae\xcb\x6b\xc1\xde\x6d\xab\x7b\xbd\xe2\x15\x43\xf2\x84\x62\x45\x21\xbb\xd3\xe5\x4c\x72\x00\x2a\xbd\x91\x1a\x9d\x09\x3f\xfc\x37\xe2\x31\xe6\x9f\xbb\x78\x8c\xfb\xd7\x86\xae\xe7\x15\x3d\xc0\xff\x15\x00\x00\xff\xff\xf2\xf5\x00\x01\xaf\x02\x00\x00")

func stringNololBytes() ([]byte, error) {
	return bindataRead(
		_stringNolol,
		"string.nolol",
	)
}

func stringNolol() (*asset, error) {
	bytes, err := stringNololBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "string.nolol", size: 687, mode: os.FileMode(438), modTime: time.Unix(1623961616, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"logic.nolol":             logicNolol,
	"math_advanced.nolol":     math_advancedNolol,
	"math_basic.nolol":        math_basicNolol,
	"math_professional.nolol": math_professionalNolol,
	"string.nolol":            stringNolol,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"logic.nolol":             &bintree{logicNolol, map[string]*bintree{}},
	"math_advanced.nolol":     &bintree{math_advancedNolol, map[string]*bintree{}},
	"math_basic.nolol":        &bintree{math_basicNolol, map[string]*bintree{}},
	"math_professional.nolol": &bintree{math_professionalNolol, map[string]*bintree{}},
	"string.nolol":            &bintree{stringNolol, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}

// Code generated by go-bindata.
// sources:
// schema/schema.graphqls
// DO NOT EDIT!

package schema

import (
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaSchemaGraphqls = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x93\x4d\xcf\xda\x30\x0c\xc7\xef\xfd\x14\xe9\x0d\xa4\x7d\x82\xdc\x78\xe9\x01\x09\x18\x23\x6c\x97\xa9\x87\xa8\x75\xdb\x48\x34\xe9\x12\x17\x09\x4d\xfb\xee\x53\xd2\xf4\x25\x05\xb4\x17\x09\x3d\xa7\xda\x4e\xfc\xcb\xdf\xae\x6d\xb2\x0a\x6a\x4e\x7e\x46\x84\x10\xf2\xa3\x05\x7d\xa7\xe4\x8b\xfd\xb8\x40\xdd\x22\x47\xa1\x24\x25\x07\x6f\x45\xbf\xa2\x08\xef\x0d\x74\x97\x7c\x9e\x54\x39\x2c\x24\xaf\x81\x12\x86\x5a\xc8\x32\x5e\x52\x72\x54\x39\x0c\xa7\x86\x92\xef\x36\x90\xba\x88\x16\x45\x41\xc9\x59\x14\x85\x73\x0d\xe8\x1b\xe8\x49\x86\x0d\x88\x6c\x86\xfc\x64\x90\xa3\x73\x39\xc2\x92\x12\xd6\x5d\x9a\x26\xd8\x57\x7c\x38\x1d\x84\xf6\xca\xbd\xd6\xbe\x24\x7f\x6f\x31\x49\x3d\x84\x47\x3b\xd9\xb4\x18\xa7\xcb\xc7\x13\x07\x07\xd9\xd6\x9d\x18\x4f\x5e\xed\x77\xdf\x12\x67\xb1\xaf\xec\x94\x6c\x2e\xce\xde\x26\xab\xad\x3f\xde\x0f\x69\x9b\x3a\xf7\x49\xec\xb2\x3a\x5f\xbc\xf5\xf9\xe4\x8c\x73\xd2\x05\xe7\x05\x4c\x65\xf9\xec\x46\x69\xa4\x64\x27\x31\x76\x6e\x56\xe7\xd4\xb2\x3b\x2f\xe8\x9e\x8b\x88\x66\xf4\x5f\xd0\x3d\xb8\x27\x39\x07\xb4\x56\xfa\x35\xe9\xf9\x5b\x33\x69\xa6\xcd\x32\x30\x86\x92\xb5\x52\x57\xe0\x72\x14\x60\xc7\xc0\xbf\x5a\x0a\x5c\x6b\x2e\xb3\x2a\x44\x95\x02\x59\xc5\xc3\xd8\x0d\xb4\x71\x83\x39\xaf\xc7\x0e\x91\xc7\xe5\x1c\xf9\x06\x24\xc2\x1f\xc4\x0b\xc3\xe0\x5a\x4c\xa4\xfd\x5d\x45\xe3\xe0\x1c\xc1\xa0\xef\x5f\xdc\x4d\xb8\x91\xbc\x61\x95\xc2\x90\x30\x9d\xe0\x59\x19\x0e\x3b\xd4\x30\xf2\xfa\xff\xa1\x64\x21\xca\x7f\xed\x79\xf0\x5c\x0f\x0f\xc1\x8f\x90\x61\x5d\xc1\xa0\x6d\x66\x9c\x06\xba\x26\xfd\x7d\x26\xea\xbd\x3d\xd7\x4d\x76\x9a\xd5\xf8\xbf\x8d\x16\x6e\x8d\x3e\x72\xb7\xe2\x77\x2e\xd7\xef\x00\x00\x00\xff\xff\xb1\x76\xd4\xad\xd8\x05\x00\x00")

func schemaSchemaGraphqlsBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaGraphqls,
		"schema/schema.graphqls",
	)
}

func schemaSchemaGraphqls() (*asset, error) {
	bytes, err := schemaSchemaGraphqlsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.graphqls", size: 1496, mode: os.FileMode(420), modTime: time.Unix(1539187019, 0)}
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
	"schema/schema.graphqls": schemaSchemaGraphqls,
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
	"schema": &bintree{nil, map[string]*bintree{
		"schema.graphqls": &bintree{schemaSchemaGraphqls, map[string]*bintree{}},
	}},
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

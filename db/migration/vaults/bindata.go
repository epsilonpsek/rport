// Code generated for package vaults by go-bindata DO NOT EDIT. (@generated)
// sources:
// 001_init.down.sql
// 001_init.up.sql
package vaults

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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x28\x4b\xcc\x29\x4d\x2d\x4e\xb0\xe6\x42\x16\x2c\x2e\x49\x2c\x29\x05\x09\x02\x02\x00\x00\xff\xff\x45\xff\xd6\x78\x2a\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 42, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x4f\x8f\x9b\x30\x10\xc5\xef\x7c\x8a\x11\x97\xec\x4a\x8b\xd4\x9e\xa3\x1e\x68\xe2\xb6\xa8\xac\xb3\x65\x8d\xba\x39\x19\x02\xd3\x16\x85\x38\xc1\xd8\x51\xf2\xed\x2b\x0c\x45\xfc\xc9\x5f\x75\x7d\x41\x32\x6f\x7e\xf3\xe4\x79\xe3\x38\xe0\x5c\x38\x96\xe3\x00\x8b\x57\x39\x42\xa9\xa4\x4e\x94\x96\x08\xbf\xb6\x12\xf6\xb1\xce\x95\x75\xad\x78\x16\x10\x97\x11\x60\xee\x67\x9f\x80\xbd\x8f\x73\x8d\xa5\x6d\x3d\x58\x00\x00\x76\x96\xda\xd0\x3d\x1e\x65\xe4\x2b\x09\x80\x2e\x18\xd0\xd0\xf7\xe1\x25\xf0\x9e\xdd\x60\x09\xdf\xc9\x12\xdc\x90\x2d\x3c\x3a\x0b\xc8\x33\xa1\xec\xa9\x06\x24\x79\x86\x42\xf1\x96\xa3\xf0\xa0\xaa\x6f\x0b\x98\x93\x2f\x6e\xe8\x33\xf8\xd0\x14\x48\x2c\x74\x26\x31\xe5\xbf\xe5\x56\xef\x6c\x60\xe4\xad\x65\x49\x8c\x15\xa6\x3c\x56\x35\x6c\x5e\xf9\xee\xb0\x06\xb2\xd5\xb1\x96\x55\x84\x13\x32\xbd\x4b\x6f\xa1\xfd\x93\x75\x69\xcd\xaf\x35\x1e\x7b\xaf\x73\xa6\x91\x79\x52\xfb\xaa\x4c\x1d\x77\x5d\xd5\x48\x66\x3d\x4e\xaf\xce\xd2\x71\xc0\xd5\x6a\x0b\x99\x48\x24\x6e\x50\x28\x30\xcd\xef\x48\x43\xf8\x62\xde\x21\x2a\x8b\x3c\x53\xc8\x4b\x2c\x34\x8a\x04\x23\xeb\x95\x30\x88\x4a\x2c\x22\xf8\x04\x1f\x8d\xe3\x9f\xdf\x48\x40\x20\x12\xf1\x06\xab\xcb\x49\x1d\x9d\xc9\x4d\x2e\x3d\x91\xe2\x01\xcb\x41\x60\x95\x09\xf1\x5d\xb1\xf5\xe8\x9c\xbc\x75\x63\x66\xac\x2d\x28\x44\xb5\x9d\x08\x1e\x46\x41\x74\x5f\x67\xe6\xee\x71\x3a\xa0\x54\x03\x3d\x57\x6f\x86\x3d\xae\x0c\xa9\xf7\x23\x6c\x01\x5a\x64\x85\x46\xde\xf6\xe2\x97\x88\x7d\x47\x4f\xa7\xdb\x0c\xd6\xb3\x54\xb1\xd2\x27\xd7\xf3\xee\xcd\x4c\x57\xbc\xa1\x8d\x96\xb2\x51\xa0\x48\x78\xf2\x07\x93\x75\x6f\x0b\x53\xec\xdd\x56\xb1\xfc\xaf\x5c\xd6\x26\xde\x2b\x98\xc3\x50\xd6\xf4\xc9\xd4\xfa\x1b\x00\x00\xff\xff\x18\x0f\x61\xdf\x44\x05\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 1348, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
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
	"001_init.down.sql": _001_initDownSql,
	"001_init.up.sql":   _001_initUpSql,
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
	"001_init.down.sql": &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":   &bintree{_001_initUpSql, map[string]*bintree{}},
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

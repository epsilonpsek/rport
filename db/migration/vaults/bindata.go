// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 001_init.down.sql (38B)
// 001_init.up.sql (1.423kB)

package vaults

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4b\xcc\x29\x4d\x2d\xb6\xe6\x42\x12\x2a\x2e\x49\x2c\x29\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x5b\xfd\x1d\xbd\x26\x00\x00\x00")

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

	info := bindataFileInfo{name: "001_init.down.sql", size: 38, mode: os.FileMode(0644), modTime: time.Unix(1620636509, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc4, 0xe3, 0x73, 0x67, 0xa9, 0xb1, 0x35, 0xc5, 0xf8, 0x5e, 0xe6, 0x48, 0xd9, 0x40, 0xfc, 0x47, 0xcd, 0x36, 0xcf, 0x97, 0x2d, 0xd, 0xe3, 0x2c, 0x4b, 0x69, 0xab, 0x4d, 0xd1, 0x70, 0x26, 0xff}}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xcf\x6f\x9b\x30\x14\xc7\xef\xfe\x2b\x9e\x7c\xc9\x2a\x95\x69\x3b\x47\x3d\xb0\xc4\xdb\xd0\xa8\xd3\x51\xa3\xb5\x27\xcb\x81\xb7\x0d\x95\x38\xc1\xd8\x55\xf3\xdf\x4f\xfc\x10\x03\xc2\x16\xa2\xf8\x12\x29\xfe\xfa\xe3\xc7\xf3\xe7\x79\x1e\x78\xff\x59\xc4\xf3\x40\xa8\x6d\x8e\x50\x5a\xe3\x12\xeb\x0c\xc2\xcf\xbd\x81\x57\xe5\x72\x4b\xce\x1d\x5e\x45\xcc\x17\x0c\x84\xff\x29\x64\x40\x5f\x55\xee\xb0\xa4\xe4\x1d\x01\x00\xa0\x59\x4a\xa1\xbf\x02\x2e\xd8\x17\x16\x01\xdf\x08\xe0\x71\x18\xc2\x43\x14\xdc\xfb\xd1\x33\x7c\x63\xcf\xe0\xc7\x62\x13\xf0\x55\xc4\xee\x19\x17\xb7\x0d\x20\xc9\x33\xd4\x56\x76\x1c\x8b\x6f\xb6\xfa\xed\x00\x6b\xf6\xd9\x8f\x43\x01\x1f\xda\x03\x06\x0b\x97\x19\x4c\xe5\x2f\xb3\x77\x07\x0a\x82\x3d\x75\x2c\x83\xca\x62\x2a\x95\x6d\x60\xeb\xaa\xee\x1e\x6b\x14\xdb\x1e\x9b\x58\x45\x98\x88\xb9\x43\x7a\x42\x1b\x6d\xf5\x09\xed\xd6\x0b\x1e\x07\x1d\xf9\x07\xbc\x6e\x23\x3d\x1b\xb3\xc7\x43\x3f\x75\x12\x23\x37\xcb\xb3\xef\xe7\x79\xe0\x3b\xbb\x87\x4c\x27\x06\x77\xa8\x2d\xd4\x97\x5f\x60\x40\xfc\x50\x77\x92\xee\x54\xa6\xe9\x7b\x5a\x16\x79\x66\x51\x96\x58\x38\xd4\x09\x52\xf2\xc8\x04\x94\x58\xc0\x1d\x7c\x24\x3f\xbe\xb2\x88\x81\x56\x3b\x84\x3b\x58\x34\xb6\x2c\x66\x15\x19\xe8\x14\xdf\xb0\x1c\x39\x6a\x6b\x6f\x2f\x32\x35\xe0\x6b\xf6\xd4\x55\xfb\x57\xb0\xba\xa5\x1b\xde\x39\x0c\x8d\xc3\x83\xd5\xf7\xd1\x7f\x5c\x75\x81\x9b\xe5\x34\xbd\x7a\xee\x59\xdc\xda\x8b\x69\x62\xcc\x83\xef\xf1\x18\xec\x74\x56\x38\x94\x5d\x39\x72\xf6\x4d\xc3\x2f\xb8\x9d\x59\xcb\x68\xcc\x4b\xab\xac\x9b\x1c\xf3\x8b\x27\x3c\xdd\xca\x96\x76\x32\xdc\x6d\x02\x75\x22\x93\xdf\x98\xbc\x0c\xa6\x39\xc5\xc1\xbf\x95\xea\x57\xb9\xde\x14\x71\xad\xec\xd0\x93\x1d\x86\xb2\x37\x17\x2c\x96\xe4\x4f\x00\x00\x00\xff\xff\xa5\xc6\x23\xcb\x8f\x05\x00\x00")

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

	info := bindataFileInfo{name: "001_init.up.sql", size: 1423, mode: os.FileMode(0644), modTime: time.Unix(1620643851, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc7, 0xe1, 0x9c, 0x90, 0x4e, 0x11, 0x46, 0x1c, 0x69, 0x5b, 0x32, 0xf, 0x4b, 0x8d, 0x23, 0x48, 0x80, 0x4d, 0x2, 0x3a, 0xe7, 0xa7, 0x6a, 0x63, 0x52, 0x59, 0xcd, 0x9b, 0xc6, 0x4c, 0x34, 0xf2}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"001_init.down.sql": {_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql": {_001_initUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}

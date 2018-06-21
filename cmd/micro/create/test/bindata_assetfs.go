// Code generated by go-bindata.
// sources:
// __tp-micro__tpl__.go
// DO NOT EDIT!

package test

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

var ___tpMicro__tpl__Go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x52\xc1\x6e\xdb\x30\x0c\x3d\x4b\x5f\x41\xf4\x94\xe4\x30\xb5\x5b\xd1\x43\x30\x0c\xf0\xda\x01\x29\x90\x6e\xc1\xd2\x9e\x55\xc2\x66\x6c\x6d\xb6\xe4\x49\xcc\xb0\xa1\xc8\xbf\x0f\x92\xec\x39\x4e\x0e\x86\xc3\xc7\xf7\xf8\xf8\x68\xa5\xa0\xc7\xf2\x27\xd6\x04\x5a\x3f\xef\xb6\x5a\x83\x09\xc0\x0d\x41\xef\xdd\x0f\x2a\x19\x98\xba\xbe\x45\x26\x79\xd1\x27\xa5\x52\xa0\x75\xb1\x7b\xd4\x7a\xf7\xb2\x8d\x4c\x4f\xb5\x09\x4c\x1e\xe2\x7f\xf0\xee\xc8\xe4\xd7\xb1\x0d\x54\xe3\x3a\xca\x6f\x1d\x72\xa3\x2a\xf3\xdb\x54\x24\xf9\x6f\x4f\x17\x1a\xc6\x32\xf9\x03\x96\x04\x6f\x52\x6c\x5c\x47\x8b\x55\x60\x7f\x2c\xf9\xed\xb4\x84\x55\x2c\x7c\xa7\x70\x6c\x59\x8a\x27\xe4\x46\x9e\xe6\x36\xf6\x9b\xb9\x8d\xfd\x66\x6e\x23\x30\xf2\xc5\xd4\x44\x99\x4d\xdd\x33\xf2\x62\x15\x9f\x85\xaf\x97\xc3\x88\xa7\x6f\x0f\x5f\xb6\x50\x7a\x42\x26\xe8\x5c\x45\xed\xa8\x93\x10\xad\x21\xdb\x8c\x02\x0f\x69\xbb\xc2\xd7\x52\xbc\x04\xf2\xa3\x02\x72\x03\xa5\xb3\xec\x5d\xdb\x92\xcf\xf4\x54\x9c\x4d\x57\x0a\x32\x1f\x1a\xb4\x55\x6c\x1c\xf4\x16\xab\xff\xba\x4b\x18\xde\x87\x2c\xf2\x80\x29\x1c\x88\x71\x83\xcf\x60\x9a\x73\x86\x4d\x3e\xef\x9d\x65\xb2\xa9\x62\x6c\x0d\x4a\x01\xd3\x9f\xa4\x96\x38\x8b\x33\x33\x85\xaf\x21\x1f\x0d\xb0\x37\x80\x71\xb7\x09\x99\x24\x23\x23\xf7\xd9\x4a\x0a\x51\xc0\xa1\x75\xc8\x77\xb7\x13\x12\x9c\x97\x42\x7c\x1e\x01\x78\xed\xd1\x63\xb7\xbe\xfa\xe8\xd1\xd6\xb4\x86\xeb\x77\xd7\x37\xeb\x9b\xeb\xf8\xfb\x74\xf5\x2a\xc5\xe9\xcc\xc5\xb0\xc1\x99\x91\x61\x47\x31\xc3\xe7\x76\x7e\x1d\x1d\x1b\xb2\x2c\x85\xb8\x9f\xec\x9c\xe4\x32\x85\x36\x9c\x19\xe2\x97\x31\x26\x9e\xd6\x4b\x11\x4c\xe8\xa8\xf8\x1c\xe2\xb9\xee\x6e\x53\x5a\xa6\xa3\xc0\xd8\xf5\x61\xb8\x40\xbc\x36\x1c\xe3\xc3\xd8\x83\xcb\x12\xa9\x36\xf1\x1f\x2b\x80\xac\x20\xc5\x57\xec\x68\x08\x5f\x8a\xa2\xa6\x04\x7c\x78\x2f\x4f\xf2\x5f\x00\x00\x00\xff\xff\x95\xda\x88\x1b\x97\x03\x00\x00")

func __tpMicro__tpl__GoBytes() ([]byte, error) {
	return bindataRead(
		___tpMicro__tpl__Go,
		"__tp-micro__tpl__.go",
	)
}

func __tpMicro__tpl__Go() (*asset, error) {
	bytes, err := __tpMicro__tpl__GoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "__tp-micro__tpl__.go", size: 919, mode: os.FileMode(420), modTime: time.Unix(1529569196, 0)}
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
	"__tp-micro__tpl__.go": __tpMicro__tpl__Go,
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
	"__tp-micro__tpl__.go": &bintree{__tpMicro__tpl__Go, map[string]*bintree{}},
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
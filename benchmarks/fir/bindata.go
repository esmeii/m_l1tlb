// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// kernels.hsaco
package fir

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

var _kernelsHsaco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4d\x6f\xdc\xc4\x1b\x7f\x3c\xf6\x7a\xfd\x77\xda\x26\x7f\xe0\x10\xca\x01\x23\x55\x0a\x42\x5d\x6b\xb3\x09\x21\xca\x85\xbc\xb1\x69\xd4\x4d\x9a\x17\x9a\x12\x50\x15\x39\x6b\xaf\x63\xc5\x6f\xf2\xda\x51\xb6\x12\xdb\x4d\x42\x21\x87\x4a\x48\x9c\xfa\x21\xf8\x00\x9c\xd8\x85\x4f\x80\x38\x71\xc8\x81\x4b\x38\x73\x29\x27\xd4\x45\xb6\x67\x76\x6d\x53\x27\x05\x05\x2a\x11\xff\xa4\xdd\xdf\xce\xf3\xcc\xf3\x36\x1e\xad\xc7\x9e\x79\xf8\x41\xa5\x8c\x28\x6a\x1a\x30\x68\xf8\x19\x28\xff\xc7\x50\xd8\x26\x8a\xe3\xb7\x42\x46\x81\x6c\x12\x38\x98\x86\x2b\xc0\x03\x0b\x00\x4c\xa4\x5f\x92\x3b\x54\x9c\x39\x2c\xa7\xb0\x5d\x1a\x9e\xf2\x71\x26\xf9\xf8\x76\xb9\x48\x3b\xc9\xbf\x53\x71\x8e\xda\xf9\xb9\x82\x80\xe5\x09\xb6\x21\xce\xc4\x0e\xfd\x45\x3b\x52\xdf\xda\xa9\x2b\x33\x2f\x60\x17\xcd\xcf\xc7\xea\xa9\x2b\xb3\x67\x8c\x4b\x1a\x18\x32\x9e\x08\xfa\x89\x47\xf8\x98\x8b\x33\x13\xb1\xe3\x70\xfc\x99\xa5\xf9\xa0\x3b\xb9\x36\x6f\x04\xf3\x21\x94\x33\x90\xef\xd5\x46\x64\x33\x4b\xf3\x0b\x2b\x77\xc3\xbe\xfe\xf4\x18\xc0\x72\xc9\x90\xd5\xaa\x59\x90\x0c\xd9\xff\xec\xd4\x25\x9f\x54\xed\x41\x55\x2f\xa8\xb5\xfd\xc9\xe2\x18\xf6\xff\x2c\x0f\xc0\x63\x9b\x42\xa1\xc0\x6f\x28\x4e\x5d\xb3\xcc\x29\x81\xe0\x13\x61\xf4\xa6\x50\x14\xee\xf3\xb7\x15\xc7\x54\xf4\x7a\x5f\xc3\x0b\x42\x41\x58\x96\x0c\xa5\x2f\x12\x04\xa1\xbc\xb8\xc6\xfb\xbc\xde\x30\xb6\x2d\x3d\xa2\x1e\x29\x2f\xae\x4d\xef\xca\x23\x81\xb6\x22\x99\xaa\x27\xa9\x7d\xd3\x3b\xb6\x62\xce\x55\x84\xb9\x98\xb6\x97\x4c\x90\x44\x49\xb8\x1f\x68\x67\x1c\xb5\x1e\x0b\xc9\x87\xf4\x9c\x64\x2c\xcf\xb5\x3d\x97\x27\xcd\x0f\x1b\xb6\x12\xeb\x33\x52\xd3\x2d\xc9\x7d\x67\xa4\xd7\x63\x5d\x7b\x10\xf7\x30\xd9\x53\xcd\xe8\x9a\x6a\x4e\x3d\x57\xb5\x21\xe9\x9e\x72\x5b\x33\x65\xa2\x5e\xd0\xad\x6d\x49\x9f\xf5\x6a\x35\xc5\x89\xf7\xf2\x53\x20\xbd\xca\x63\xa5\xbe\x77\x59\x76\xd6\x6d\xa9\xaa\xac\x7a\x92\x3e\xd5\x73\xd1\xd7\x57\xab\x44\x13\x62\x5e\xa9\x49\x9e\xee\xa6\xd7\x5e\xb5\x94\x5a\xed\x72\x96\xae\x99\x97\xf6\xaa\xef\x68\x75\xd7\x72\x1a\x97\xb3\x78\xd3\x33\xb6\x5c\xc9\x4e\x2f\xde\xd3\x4c\x37\xbd\xf0\xf1\xf4\xc2\xc7\xd3\x0b\x9f\x6d\x04\xa2\xf4\x9a\xef\x46\x6b\x3e\xaf\xa6\x0b\xba\x1a\xb7\x34\x59\x56\xcc\x70\x40\xef\xd4\x6a\x75\xc5\xfd\x28\x3d\xc1\xc5\x89\xf1\x7f\x3e\xfe\xe6\x4b\x8e\xff\xf1\xf9\xf1\xe7\x2c\x59\x59\x71\x2c\xbb\x77\x77\xc1\x26\xfe\x9d\x4f\x72\xd4\x75\x45\x35\x14\xd3\x0d\x33\x7c\x8f\x5c\xd3\x05\xc7\xf2\x6c\xac\x2a\x6b\xfb\x8a\x1c\xea\x8b\x58\xbd\xe2\x68\x7b\x92\xab\xa4\x77\x88\x3b\xc7\x35\x92\xca\xee\x49\x7b\x4a\xcd\xb1\x48\x50\x41\xe8\x8d\xd4\xb2\x67\xac\x2f\xac\xac\xf5\x6f\x84\xa3\x63\x7d\xcd\x46\x5c\x33\x8a\x35\x4b\xd2\x7e\x59\x97\xdc\x7b\x96\xb3\x1b\x66\x1d\x38\x2d\xbd\x3b\xc1\x8b\xa2\xc8\xa7\xaf\x6b\xfc\xb5\xc8\x75\x7f\x05\x96\x58\xdf\x51\x91\x0f\x5e\x2f\xa1\x69\xdc\x56\x7f\xf9\x62\x00\xe1\xa5\x0f\x95\x74\x58\x5e\x5c\xfb\x1b\xcb\xab\x0c\x19\x32\x64\xf8\xcf\xa1\xff\x3f\xca\x85\x4f\x76\x7f\xfa\xc3\x8c\x63\x16\xbe\x86\x1f\x82\x67\xbd\xf8\x9f\xf6\xad\xc8\xef\xab\xc1\x93\x60\x1f\x0c\xc3\xb0\xdd\x6e\xb7\x7b\x91\x79\x5f\x14\xbe\x07\xd4\x61\x82\xca\xd9\x8e\xdf\x3e\x02\xd4\xf1\x1f\xd1\x69\x8a\xed\xbc\x0d\x00\x2d\x84\x9a\x0f\xe1\x71\x9b\xee\xd2\x9f\xfb\x15\x70\x34\xf7\x15\x07\x50\x62\x00\x4a\x2d\x38\xfa\x16\xb5\xd8\x76\x0b\x31\xcd\x1b\xf0\xa8\xed\xdb\x04\xa3\x88\xd8\x8e\x7f\xbb\x3a\xa2\xd8\xce\x70\xe0\x83\x6b\x02\xcd\x37\xfd\x7e\x34\x3c\xf9\x14\x21\xb6\xd9\x62\x59\xf0\xfd\xfe\x0f\x5d\x69\xe6\xd1\xb5\x26\x87\x86\x9a\x2c\x7a\xa5\xc9\x0e\xbe\x06\xdc\xab\x43\x40\x0f\xb0\xa5\xef\x78\xbe\xc4\xc2\x97\x3f\x1e\xb2\x08\xb8\xab\x57\x4a\xfc\xe0\xb5\x49\x80\x95\x13\x16\x80\x3e\x44\x28\xc8\x97\x3e\xa0\x0f\x98\xcf\x98\x56\xae\x95\x3b\x44\x74\xbe\x6d\xc3\xe3\x36\xcf\x32\x37\x4f\xbb\x8f\xda\xe0\xdb\x02\x02\x0a\xb1\x4d\x00\x28\xd1\x34\x9a\x04\xb0\x4f\xc2\x77\x02\x07\xed\x97\x3d\xf6\x19\x32\x64\xc8\x90\x21\x43\x86\x0c\x19\x32\x64\xc8\x90\xe1\xdf\x03\xd9\x6b\x3e\xc6\x8f\xf3\xe4\xc9\x7d\x18\x73\x0e\xf3\xaf\x58\xcf\x27\xe4\x4f\x9f\x75\x2d\x9f\xbf\xc1\x0a\xb2\xaf\xfc\x53\xca\x3b\xdd\x8a\x66\xee\x2a\xce\x94\x50\xa9\xcc\x0b\xe3\x62\x11\xaa\xba\x64\xaa\xc2\x5e\xb8\xd9\xea\x4b\x84\xb3\xf3\xa5\x70\xd6\xf6\xff\xe3\xf2\x3c\x96\x9f\x24\xe4\x41\x13\xe5\xfb\xfb\xef\x18\x57\x53\xde\x2b\x83\x68\x5a\xae\x02\xa2\xdc\x30\xeb\x0d\x03\x44\xd5\xf4\xc4\x1d\xa9\xbe\x03\xf8\xdb\x97\xbb\x0e\x88\xae\xb2\xef\x06\x2d\xc9\xd0\xaa\x20\x56\x2d\xc3\x50\x4c\x17\xc4\x7a\xc3\x70\xa5\x6d\x10\xeb\x3b\x75\xd7\x09\x7f\x85\x0c\xb3\xb3\xc5\xad\x52\xf0\x3d\x16\xbc\x84\xde\x9a\xdf\x5c\x9e\x59\x5a\x9c\x3b\xbb\xde\x17\x05\x85\xc7\x80\x6c\xfb\xa7\xed\xff\x13\x24\xcf\x5e\xe4\xf1\xb5\x27\x66\x64\x3e\x10\x2e\x62\x79\x0e\xc7\x22\xee\xc8\x3c\x19\x04\x80\xdf\xba\x5d\x8b\xd8\x93\xf9\x40\x58\x48\xa4\x95\x48\x07\x5e\xc7\xbe\x89\x9e\xcc\x1f\xc2\xc3\x09\x7b\x26\xc1\x6f\xe2\x73\x09\x44\x4f\xe6\x2b\xe1\x5c\x22\x5e\xf2\x35\xd7\x48\x7c\x2f\x21\xf5\x5c\x49\x9a\x83\x02\xb6\xa5\x89\x20\xe5\xbc\x47\x2e\x51\x3f\x09\x33\x81\x5d\x16\x13\x61\x6c\x6c\x7f\x23\x25\x3c\xe1\xf7\xa3\xd7\x3e\x82\x27\xd8\x7e\x1f\xb7\x07\x22\x67\x3e\x20\x32\xae\x0b\xd1\xdc\x23\x18\xc2\xe7\x7e\x36\x53\xe2\x13\xac\xa6\xd8\x6b\xd8\xfe\xfa\x39\xf6\x7f\x04\x00\x00\xff\xff\x01\x53\xf2\x07\x88\x24\x00\x00")

func kernelsHsacoBytes() ([]byte, error) {
	return bindataRead(
		_kernelsHsaco,
		"kernels.hsaco",
	)
}

func kernelsHsaco() (*asset, error) {
	bytes, err := kernelsHsacoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kernels.hsaco", size: 9352, mode: os.FileMode(509), modTime: time.Unix(1536376204, 0)}
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
	"kernels.hsaco": kernelsHsaco,
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
	"kernels.hsaco": &bintree{kernelsHsaco, map[string]*bintree{}},
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

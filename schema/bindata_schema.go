// Code generated by go-bindata.
// sources:
// schema.json
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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1b\x4d\xaf\xdb\xb8\xf1\xfe\x7e\x05\xe1\x2d\xb0\x17\xd9\x6f\xd3\x2e\x72\xd8\x3d\x05\xd9\xb6\x28\xd0\x6e\x0b\x6c\xf7\xb4\xc8\x81\x92\xc6\x16\xf3\x24\x52\x25\xa9\xe7\xb8\x41\x7e\x54\xff\x48\x7f\x53\x87\x94\x48\xcb\x36\x25\x51\xcf\x1f\x71\x8a\xe8\x90\xf8\x51\xc3\x99\xe1\x70\xbe\x49\x7d\x7c\x20\xf8\x2c\x7e\xa7\xb2\x02\x2a\xba\xf8\x81\x2c\x0a\xad\xeb\x1f\x1e\x1f\xdf\x2b\xc1\x97\xed\xe8\x4a\xc8\xcd\x63\x2e\xe9\x5a\x2f\xbf\xfb\xfe\xb1\x1d\xfb\x66\x91\xb4\x33\x73\x50\x99\x64\xb5\x66\x82\x9b\xd9\x6f\x45\x55\x97\xa0\x81\xb4\x60\x64\x2d\x24\xa9\x59\x0d\xcf\xec\xdf\xa4\x02\xa5\xe8\x06\xd4\xca\xcd\xd5\xbb\x1a\xcc\x24\x91\xbe\x87\x4c\xbb\xd1\x5a\x8a\x1a\xa4\x66\xa0\xf0\x5d\xcb\x9f\x1d\x07\xfe\xcc\xa4\xe0\x15\x70\x7d\xf8\xe6\x00\x17\x95\x92\xee\x3a\x54\xfe\x65\xc5\xf8\x5f\x34\x54\x66\xda\xab\xa3\x57\xac\x1b\xff\x88\x42\x90\xb0\x36\x28\xbe\x79\xcc\x61\xcd\x38\x33\x6b\x52\x8f\x3d\xb2\x0b\xf2\xc9\x4f\xfe\xb4\xc7\xb3\x28\xc5\x86\x65\x4b\xa5\xa9\x86\x5b\x72\xd6\x23\x3b\xc4\x59\x4e\x35\x55\x70\x53\x79\x75\x24\x87\x38\xc2\xdd\xcd\x50\x0f\x6e\x2a\xa8\x8e\xe6\x10\x4b\x99\xa8\x2a\x76\x53\x19\xb5\x14\xc7\xf9\x59\x56\xa0\xe9\xcd\x79\x6a\xa9\x0e\x30\xb6\x6b\xaa\x65\xfd\xb4\xb9\x2e\x53\x7d\x9e\x1c\xc1\x3e\x3f\x0f\x3d\xae\x16\x34\xcf\x2d\x28\x2d\xff\xd1\xf7\x1a\x6b\x5a\x2a\xf0\x0e\xca\xe3\x1b\xf4\x27\x81\x15\x31\x5d\xda\x25\xfd\xb1\x07\x75\xc4\xfd\x91\xeb\x7b\x43\x24\xd4\x12\x14\x42\x52\x33\x46\xc4\x9a\x50\x4e\x7a\x64\xc8\x92\xd4\xc5\x4e\xb1\x8c\x96\x09\xc1\x51\xdd\x98\x1f\xe8\x21\x33\x81\x73\x18\x07\xb9\x3a\xa6\x11\x74\x91\xfe\xed\x80\xab\x3c\x9e\x7d\xfa\xa6\x13\x40\x53\xe1\xdb\xdf\x10\x4f\xc7\xd5\x22\x21\x8b\x8e\x2f\xf3\xd3\xb3\xb5\x20\xef\x92\x30\x0e\x94\x2e\x6d\x4a\x23\x40\x3f\xf1\x04\xf0\xd3\xe9\xdc\x85\x08\x73\x7c\xc4\xd7\x96\xf1\x5c\x6c\x95\xe1\xa5\x64\xbc\xf9\x60\x7e\xe4\x54\xe2\xb0\xf9\xb5\x96\x00\xa9\xca\xcd\xcf\x86\xb3\x0f\x51\x4c\x5a\xc0\x28\x0e\x51\xb5\xa4\x71\x1a\x83\x6c\x86\xcd\xc8\x4d\x8b\xa2\xc1\x59\xf6\x34\x4c\xc0\xed\xbd\xd2\x92\xf1\xcd\x62\x70\x71\x47\x5a\x68\x90\x72\x5a\x01\x61\x39\xea\x1c\x5b\xef\x70\x32\xd1\x05\x53\x7d\x55\x5c\x91\x9f\x3b\x30\x45\x32\xd4\xd2\x14\x48\xa3\x20\x27\x54\xa1\x16\xaf\x41\x9a\x20\x6b\x83\x77\xbb\x36\x8b\x01\x48\xc1\x40\x52\x99\x15\x46\x57\x10\xae\xb4\x7a\xae\x0a\x56\xe3\x7c\xbd\x05\xe0\xc7\x0a\x4f\x79\x4e\xd0\xb5\x5a\x05\x37\x28\x13\x92\x36\x9a\x70\xa1\x2d\x6e\x09\x88\x86\xe3\x4c\x21\x9f\x9c\xe0\x68\x5a\xc2\x01\x6a\x85\xac\x02\x72\x66\x26\x65\x42\x22\x4c\x2d\x10\xab\x16\x88\x7c\xd7\xa2\xb0\x21\xf0\x60\x45\x62\x09\x1f\x98\xd2\x84\x21\x47\x64\x53\x8a\xd4\x10\x32\xaf\x6a\x9a\x81\xa5\x4d\xcb\xb2\xcf\xa9\x22\x4f\x5c\x6c\xb9\x43\xeb\x52\x16\xc6\x11\x37\xcf\x20\x21\x4a\xe0\x60\xf6\x64\xc4\x50\x91\x2d\x53\x50\xee\x56\x71\x9b\x8c\x46\xfa\x8c\x5b\x21\xa3\x37\x3a\x0a\xeb\x68\xce\x71\x82\x39\xe4\x98\x3d\xd0\xb0\x83\xf6\x20\x2f\xcd\x45\x26\xe4\x32\x90\x09\xdc\x9e\xfd\x40\x86\x30\xc6\xfa\x60\x5a\x75\x7b\xce\x03\xe9\x96\xe7\xfc\x61\x64\x1d\x11\xa1\x73\x0f\xca\x77\x7f\x37\xa4\x7f\x3b\x21\x81\x6c\x49\xf8\x57\xc3\x24\xe4\xad\xc7\xb6\x1e\x8d\xbc\x0b\x49\xed\x04\xd6\xb9\x4a\x03\x7e\x00\xfd\x2e\x98\x7d\x0c\x39\xe4\x33\x63\x64\x21\x94\x36\xae\xe1\x5c\x3f\x8c\x4e\xa5\xa2\x36\xc6\x18\x8c\x4b\x8b\x32\x4a\x99\x58\xfd\xfc\xfd\xe5\x88\x5b\x6c\xb1\x74\x5f\x5f\x94\xee\xeb\x00\xdd\x71\x15\x8c\xd7\x2b\xbf\x4b\x91\xba\x65\xc5\x10\x0f\xfb\x3a\xa0\x84\x33\xec\x25\xa8\xaf\x18\x5d\x96\x98\xb5\x9c\x06\xf8\x09\x85\x3d\x0a\xe8\x3f\xd9\x3f\x53\x8c\x6a\x94\x18\x74\x24\xa5\x18\x86\xd6\x52\x54\x38\x80\x61\xdd\x84\x66\x61\xa3\x73\x3f\xf0\xea\x82\x6a\x97\x56\x2a\x0c\xc1\x2b\xf2\xed\xdb\xee\xaf\x6f\x6d\x54\x54\x3e\x15\x45\x34\x98\xb3\x2a\xb0\x63\x42\x29\x86\x11\x78\x45\xfe\x69\x12\x06\x5a\xa3\xf5\xd0\xac\xc0\x20\x0d\xb9\x65\x00\x83\x30\xa6\xb5\xba\x10\xcd\xa6\xd0\x27\xd9\xea\x94\xad\xf5\x4c\x78\x2a\x7f\x9a\x50\x9d\x21\x4f\x72\xee\xa6\xf5\xe3\xd8\x85\x1d\x4d\x6c\x32\x9e\x32\x4e\xe5\xae\xcd\xbf\x73\x68\x73\xdf\x54\x9a\xa1\x9e\x63\x1c\x90\x4b\xcb\x08\xd5\x45\x2b\xe2\x23\x23\x1e\xcd\x41\xe3\xa0\xcb\x8d\x14\x4d\x1d\x0f\x7f\x54\x64\x8d\x34\x59\x5a\x5b\x09\x13\xb5\xeb\x1f\x12\x6c\x5f\xb8\x97\x8c\xb7\x27\xcb\x8b\xf3\xaa\xf9\x34\x93\x41\xfd\xf1\x50\x82\xc3\x80\x57\x74\x4f\x18\xbd\x47\x30\xa1\x88\x27\xf0\x5d\x37\x62\xbf\x68\x32\xba\xea\xfe\x13\x90\xc0\x01\xea\x43\x33\x75\x6d\x8f\x81\x3a\xcd\xcf\x8a\xb2\xda\x19\xac\x5c\x58\x5e\xcf\x20\x55\xeb\x9c\xa3\xb4\x24\x92\x49\x8b\xfa\x50\x5e\x8e\xd0\x97\x2e\x30\x05\xd5\x73\x5b\xff\x5c\x57\x5e\x1d\x9d\x2b\x89\x2b\xf8\x26\xce\x23\xdf\x51\xbd\x80\x39\x01\x5f\x1a\x7e\xe6\x57\x0c\x87\xc2\xb6\x51\xe6\xfc\x80\xbb\xe7\x67\x66\xb8\x1d\x76\x93\x03\x22\x8e\xd4\xda\x85\x2b\x06\xa2\x02\xdc\x31\xb3\x1f\xfb\x91\xbc\x82\x9c\x61\x22\x91\x0f\x24\xa2\x7e\xae\x6a\x52\x53\xc2\xcd\xa6\xc9\xb8\x06\x49\xb3\x2e\x53\x9c\x30\xda\x3d\x5b\x72\x6b\x92\x0a\x29\x42\xf9\x84\x7b\xc6\xc8\x9a\x2d\xfb\x19\x46\xe3\xb9\xdd\x55\x1e\x2e\x4b\x47\xd0\x1f\x27\xe6\xbd\xf5\x25\x9d\x68\x92\x3d\xfd\xc1\x7e\xdf\x2c\xfb\x0e\x55\x08\x5f\x15\x28\x00\x7c\x59\x05\xfa\xd5\x34\x62\xa7\x34\xa8\x6d\xeb\x5e\x49\x85\x7e\x1d\xed\x19\xff\xdf\xe8\xd0\x9a\x95\x43\x55\xb3\x9f\xf7\x55\x7f\xf6\xc8\x3f\x8b\xfe\x44\xb4\xbf\xbc\x4b\xbd\x44\x3b\xa1\xd6\x42\x76\xe7\x53\xa2\xd1\x1b\x61\xba\x08\x5d\x23\xde\x74\x0d\x38\xd8\x15\xff\x68\x3a\x0c\x54\x6e\x40\xfb\xb2\xfc\x56\x41\x78\xb2\x2b\x77\xb2\xfa\xd1\x46\x95\x87\x8e\xea\xd2\xb9\x67\x4c\x07\x6b\x21\x8f\x0c\xc6\xa8\xca\xc6\xa4\x9f\xa3\xd3\xa4\xd0\x62\x4e\x26\x1c\xa7\xa6\x5e\x5c\x49\xc7\x59\xe2\x48\xdd\xab\x83\x1b\xed\x7a\x7a\xa8\x97\x6e\x6e\xb8\x0b\x3a\xb2\x9a\x3d\xff\x77\xb5\xaf\x76\x1d\x5f\xd2\x9e\x0e\x77\x94\x3d\xd4\x19\x7b\x1a\xe8\x30\x8f\xac\x66\xcf\xff\xbd\xed\xe9\xeb\xab\xed\x69\x6c\x20\x69\xba\xf0\x79\x95\x48\x52\x0a\xd3\x56\x36\x24\xc2\xc1\xc4\x96\x8e\x33\xbb\xa7\x23\x4d\xcd\x87\xb0\x04\x02\xb7\x8e\x2e\xdd\xd2\xfd\x6c\x87\xb0\x2f\xeb\x4d\x36\x6a\xb0\x0b\x73\xad\xd6\xed\xcc\x76\x71\xb6\xcd\xe3\x81\xdd\xd1\x67\xf4\x84\x9a\xe5\xf1\x6e\x60\x51\x32\xa5\x61\x38\xaf\xbd\x5d\xcb\xc6\xe4\x5e\xcb\x8e\x9b\x33\x9b\x36\x07\x0a\x9b\xb4\x12\x39\xbf\x87\xd3\x67\xf0\x52\x87\x5d\xe6\x40\xab\x45\x09\x92\x28\x96\x43\xe7\x64\x6c\x42\xbe\xec\xec\xb9\xe7\x5e\x56\xe4\x6d\x23\xcd\xb5\x95\x72\x47\xde\x37\x4a\x13\xe3\x5e\x95\xbd\x80\xa2\x44\xf6\x04\xe6\xee\x49\x29\x70\x04\xd1\x94\xe2\x19\xec\x91\x16\xe4\x27\x47\x58\x57\xcb\x68\x43\xb5\x99\x0d\x01\x13\xb5\xd9\xb9\xc1\x2b\x2e\x18\x8f\x29\xb1\x07\x9e\x56\x66\x0f\x3a\xcb\x57\xb9\xe7\x9c\xa8\xea\x2a\xb3\x2f\x23\x53\x0a\x29\x43\x77\x87\x6d\x7c\x57\x23\xa3\xe0\x08\xe7\x16\xcf\x80\xec\x42\x2d\x5d\x3f\xe7\xe2\x89\x48\xcf\x7f\x9f\xe1\x32\xde\x90\x0e\x8f\x3d\xb3\xe6\x4f\xed\xf9\x34\x10\x74\x13\x95\xbd\x1e\x56\xee\x92\xee\x7f\x92\x4a\x41\xf3\x72\xb7\x22\xcf\x54\x32\xca\x35\x79\x45\xcc\x29\x37\x51\x50\xae\x97\x4a\x34\x32\x83\xdc\xa1\xfb\xd1\x03\xfd\xbe\x05\x5a\x37\xdc\x7a\x1a\x93\xe1\x9c\xc0\xfc\xa1\x85\xe9\xc6\xf7\x17\x3a\xef\xbc\x6b\x9d\xa1\x60\x34\x2c\x35\x73\x93\xf7\xe9\x76\xee\x5f\x8c\x62\xd8\xa0\x83\x56\x4c\x1d\x2b\xf3\x7f\xff\x73\x7a\xa3\xc2\x3d\x71\x4a\xe9\xca\x59\x47\xe0\x5e\x4d\xf9\x0e\x04\x1f\xe7\xe2\x47\x4f\x9d\xe7\x2e\xdb\xc3\xcf\xba\xd2\x31\x4e\xdb\xe2\x1b\x72\x0a\x93\xcb\x8b\x89\x60\x7e\x52\x7c\x24\xf3\x53\x5e\x14\xd1\xdc\x13\xb3\x72\xc5\x69\x3d\xa9\x10\xa3\x68\xe6\x9d\x96\xfa\x96\xde\x5e\xe8\x49\x9f\x8b\x5b\x1e\xa3\x7e\x75\x09\x97\xa9\xb7\xf6\xac\xce\xb9\x03\xe4\x67\xb5\xe7\x00\x33\x7c\xca\x15\xd3\xc6\xd9\xd7\x5d\xdd\x73\xb6\x86\xdd\xb2\x15\xe3\x6e\xde\x5c\xa4\x66\xc2\x44\xc6\xe6\x30\x36\xfd\x90\xa2\x24\x2d\x7e\xd2\xa2\x59\x91\x3f\x09\x49\xb8\xd8\x26\x6d\x81\xb4\x31\xb7\x02\x71\x21\x29\x4d\x31\x35\x6a\xef\xc4\x6f\xe9\x0e\xeb\x24\xc1\x71\xe0\x97\xb7\x7f\x53\xed\x35\xc2\x82\x62\xb9\x44\xb5\xa8\x58\xd6\x61\x44\x9f\x4e\x7e\x7a\xf3\xe7\xd9\xf7\xff\x54\x41\x5f\xc5\x97\xeb\x12\x6a\xa1\x98\x16\x72\x37\xab\x27\x10\x6f\x7a\x0b\xda\xe8\x42\xcc\x68\x89\xa0\x7d\xd8\x0d\x89\xef\x38\x50\x19\xfc\x5c\xf1\x64\xa7\x63\x9b\x08\xdf\x5d\xa4\x37\x34\xa3\x63\x60\xb7\x2c\xe9\x04\x9b\x78\x91\x25\x7b\x61\x24\xfb\x65\x26\x07\x9b\x76\x89\xcb\x21\x11\x5f\xc2\xcd\x34\x12\x83\xcc\x78\x11\x42\x53\xd1\x68\xb4\x99\xce\x48\x96\x58\x1d\x50\x9e\x15\x08\x82\xa6\xa0\xe9\x06\xed\x49\x29\x91\xd9\x03\x78\xb2\x65\xba\x20\x4c\x27\x58\x56\x28\x6d\x2a\x80\xb6\x87\x92\x10\xd0\xd9\x95\xcd\xc0\x50\xfc\x25\x78\x21\xd5\x83\xf4\x3a\x0a\xc8\x33\xd8\x4f\xa2\x6a\xe0\xb9\xed\xaf\x63\x2a\x41\x59\x69\x2f\x11\x44\x75\x09\xcd\xd2\xbf\x80\x66\xa6\xdb\xad\x7b\x63\x75\xae\x6d\x9d\x6f\x25\x93\x9f\x65\x46\xd6\xd1\x35\xcd\x9e\xe8\x06\x12\x73\x11\xbc\xe1\x39\x48\xa5\x85\xc8\x49\xba\xb3\xf5\x34\x52\x71\x10\xa4\xa2\x1c\xff\x93\xed\x27\x63\x8c\x13\x59\x57\xcb\x94\x9a\x3f\xfe\x6a\x3e\xd3\x23\x39\x33\x02\x4a\x1b\x1b\xb3\x67\x9b\xc7\xac\xf4\x69\xe2\x36\x66\xa8\xaf\x5d\x8b\xac\x88\xef\xa8\xe1\x9e\x95\x80\x6b\x9b\x11\x55\xe4\x31\xfe\x81\xa3\x8a\x51\xe5\x70\xc9\xae\x5b\x5f\xb2\xe7\x24\x71\x8b\x48\x3a\x62\x2f\xd4\xa1\x87\xf6\xdf\x4f\x0f\xff\x0b\x00\x00\xff\xff\xd1\x0a\x6e\x87\x06\x40\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 16390, mode: os.FileMode(420), modTime: time.Unix(1445009849, 0)}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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


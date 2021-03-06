package fixtures

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var __1_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x91\xcd\x8a\xc2\x30\x10\xc7\xef\x7d\x8a\x92\xf3\xee\xd2\xb2\xbb\x08\xbe\x8a\x78\x18\x93\xa1\x1d\x6c\x93\x98\xa4\xd5\x52\xfa\xee\x26\x11\xd4\xd6\x4b\x40\xd4\xc9\x65\xc8\xff\x23\xfc\xc8\x98\xe5\x7e\x18\xca\x9e\x8c\x92\x2d\x4a\x67\xd9\x3a\xdf\xc4\xdb\x30\xe3\x75\x8b\x3e\x37\x68\xf4\x3a\xd3\xf5\x60\x89\x43\xc3\xbe\xe6\xba\x0a\x69\xd6\x49\x3a\x2d\x15\x10\xc2\xa0\x0d\xf2\xbc\x32\x8a\xb5\xb2\x4e\x42\x7b\xa9\x36\x4a\x7c\x1f\x71\x57\x94\x8b\x8a\xe8\x24\xdd\xff\x05\x57\x59\xfc\x94\xe1\xb0\x99\x65\x5a\x3c\xea\xbb\x7a\x12\x68\x42\xc2\x00\xdf\x5b\x0d\x1c\x6f\x91\x3b\xfb\x87\x38\x45\x1a\xe6\xef\x8b\x31\xfd\xe7\xbb\xee\x91\x32\x99\xc5\x3a\xa8\x30\x81\xe3\x3f\x9d\x43\x50\x45\x0e\x1a\xc5\x11\xe4\x3b\x51\x0e\x90\xc0\xb1\x7a\x9a\x23\x6e\xdb\x6c\xca\xce\x01\x00\x00\xff\xff\x3f\xf8\xc5\x3d\x81\x03\x00\x00")

func _1_json_bytes() ([]byte, error) {
	return bindata_read(
		__1_json,
		"1.json",
	)
}

func _1_json() (*asset, error) {
	bytes, err := _1_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "1.json", size: 897, mode: os.FileMode(420), modTime: time.Unix(1431118616, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var __2_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x56\xcd\x92\xda\x30\x0c\xbe\xf3\x14\x4c\xce\xa5\xc4\x71\xec\xd8\x7d\x88\xde\x7a\xea\x74\x3a\xb6\xa4\xb0\x99\x26\x71\x36\x31\xd0\xed\x0e\xef\x5e\x27\x0c\xcc\x00\x81\xbd\x64\x4a\xb7\x45\x07\x7e\x22\x59\xfa\x3e\xe9\x53\xc6\xaf\xb3\x79\xb0\xa8\x74\xab\x02\x16\x9d\x37\x9e\xba\xe8\xd3\xfc\xeb\xf0\xb4\xb7\xd7\xe3\xaf\x21\xce\xbf\x34\x14\xfc\x11\x38\xa4\xe8\xc3\xa9\xaf\xc0\xe0\x39\x8d\x1f\x9e\x83\xab\xaa\xc2\xf7\xa7\x30\x05\x25\x74\xac\x6c\xac\x78\x0e\x64\x63\x54\x4a\x0b\xe0\x12\x6c\x6a\x53\x11\x03\x58\xa6\x88\x43\x74\x92\x64\x77\x56\xa7\x2e\xe0\x47\x9f\xcd\xad\x5b\xd3\x34\xe7\x28\xca\x55\xeb\xd6\x4d\xef\x6f\x5a\x87\xe7\xde\xc6\xf8\xa7\xde\xb7\xdc\x98\x76\xb9\xdd\x6e\x97\x23\x19\xa8\xde\x14\xad\xab\x2b\xaa\xfd\x38\x21\x83\xd8\x52\xd7\x8d\x3a\x87\x80\x27\xd7\xf9\xda\x54\x74\x40\xb1\xd8\x06\xae\x2c\xba\x08\xde\xdd\xe4\x89\xc6\x9b\x8e\xfc\xe9\x3c\x0e\x76\xa5\xf2\xa1\x2a\x52\x6e\xd6\xa5\x9f\xa3\x3d\x63\x77\x8c\x3c\x4c\xb2\x22\x2c\xc2\xd4\xcf\x3b\x75\x8c\x03\x57\xd7\x9f\x69\xbc\x13\xd7\x09\x63\xcf\x77\x3c\xe3\x70\xa0\x71\x6d\x9f\x92\xf3\x58\xde\x8a\x6a\x9d\x77\x7d\x4e\x0f\xcd\x65\xf7\x7a\xdb\x5d\x41\x5d\xd4\x9e\x5a\x03\xbe\x70\x75\x7f\xbe\xdd\x5e\xa3\xd7\xad\x6d\x37\xb0\x8b\x82\x12\xbe\x87\x76\xbd\x31\xa4\x6f\xb3\x91\xd2\x93\x2e\x09\x93\xc6\xf0\xc4\xe4\x2c\x93\xca\xa4\x4a\xc4\xd6\x64\x94\x84\x2f\x14\x56\x31\x84\x9c\xc0\x48\x9d\xb3\x69\x96\x24\xac\xfc\xea\x02\xe4\x9f\xdf\x92\x3d\x8c\x77\xbb\x20\x5f\xea\xe2\xe7\xed\x0d\x39\x69\x69\x59\xd8\x65\xf5\xd2\x3d\x97\xfb\xcf\x8f\x9d\x0b\xc3\xfa\x5f\xf4\xcd\xf3\xcc\x22\x65\x09\x62\xc6\xac\xe2\x9c\x67\x8a\x48\x71\x0d\x20\x35\xe3\x98\xa6\x22\x4b\x6d\xa6\xf9\x34\xfa\x7e\x36\xf7\x17\x77\xc0\xf0\x50\xf6\x45\x03\xfe\x3d\x65\xe7\x5c\x5a\x02\xcb\x33\xcb\xb4\x40\x80\x0c\x2d\x31\xcd\x8c\x10\x06\xb8\x15\x56\x0a\x92\x69\xce\xf4\x44\x6f\x6e\xb4\x0b\xa4\xcd\xfd\xe5\x1d\x40\xbc\xe7\xcb\xcd\x43\xe1\x47\x60\x6f\x2a\x5c\x28\x10\x3c\xd1\x4a\x4b\x80\xc4\xc8\x84\xe7\x0c\xc3\x1d\x9e\x61\x78\x7f\x87\x3b\x89\x8e\x65\x10\x3f\x68\x33\x8d\xc2\xcb\x5f\x29\x63\x7f\x8f\xc6\x93\x87\xc6\x2f\xec\x9e\x1a\x9f\xed\xff\xef\x66\xbf\x03\x00\x00\xff\xff\x6a\xf8\x75\x1d\xc4\x0e\x00\x00")

func _2_json_bytes() ([]byte, error) {
	return bindata_read(
		__2_json,
		"2.json",
	)
}

func _2_json() (*asset, error) {
	bytes, err := _2_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "2.json", size: 3780, mode: os.FileMode(420), modTime: time.Unix(1431118616, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var __3_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x9a\xdf\x8f\xdb\xc6\x11\xc7\xdf\xfd\x57\x2c\xf4\xe2\x16\xb0\xcf\xfb\xfb\x07\x5b\x14\x75\x02\xa4\x68\x80\x34\x01\x2e\x41\x1f\x0a\x3f\xcc\xee\xce\x4a\xb4\x25\x52\x10\x29\x5f\xae\x41\xfe\xf7\xce\xf2\x7c\x86\xef\xdc\xc2\x94\xce\x14\x7a\x96\x61\x59\xd2\x1e\x29\x7e\xf6\x3b\xf3\x9d\x19\xfe\xf6\x8c\xd1\xcf\x2a\xf5\xbb\x5d\x3b\x0e\xab\x86\xfd\x6b\x7a\xa1\xfe\xfc\xf6\xf1\xd9\xf4\x91\x61\x03\x82\xde\x5f\x19\x9f\x8c\x92\xc1\x07\x9b\x92\x04\x2b\x55\x11\x39\x7b\x2f\xb2\x57\xc1\x86\x12\xb8\x55\x2e\xa6\x00\xab\x17\x0f\xd7\x1f\x70\xdf\x0f\xed\xd8\x1f\x6e\xeb\x6f\xd9\x8c\xe3\x7e\x68\x5e\xbd\x5a\xb7\xe3\xe6\x18\xaf\xe8\xf8\xaf\x86\x1c\xfb\x5b\x3c\xbc\xda\xb7\x7b\x7c\xdf\xfe\xfb\xf1\xfa\x0c\x23\xae\x18\x2d\xfd\x79\x73\x64\xdf\x43\xc7\x64\x60\x82\x37\xd2\x37\x46\x31\xc9\x85\x61\x2f\xb9\xe1\xfc\xf1\x32\x38\x8e\x9b\xfe\x30\x2d\xbc\x86\x1d\xfb\xa6\x1e\x82\xfd\x79\xc4\xb4\xf9\xeb\x00\xbb\xe9\x88\x57\xfd\x61\xfd\x97\xc7\xeb\x86\x63\x7c\x8b\x69\x9c\x16\xfe\xb4\xe9\xbb\x5b\x76\x77\x8d\xd8\xd8\xb3\xa1\xdd\x1d\xb7\x74\x3a\xac\xf4\x87\x77\x0c\xe8\xb5\x76\xff\x78\xfd\x1e\x0e\xd8\xd1\x15\x65\x9f\x5e\xd2\x8f\xef\xa6\x84\xd2\xd9\xac\x6d\xb4\x11\x84\x14\xc2\x67\xad\x9c\x48\xce\x25\x59\x6c\xf6\x5a\x4a\x15\x4a\x96\xab\x07\x4b\xdf\x7c\xfc\xdf\xef\x2f\xbe\x84\x49\x70\x67\x39\x0f\x89\x4b\xce\x8d\x96\xe8\x40\x15\x6f\xc1\x85\x98\x93\x48\x26\x47\x91\xb4\xf6\x52\x2f\x86\xe9\x9f\x98\xd9\x35\xee\x99\x70\x8c\x8b\x46\xd1\x63\xc2\xa4\x09\x93\x5e\x06\xd3\xeb\x9c\x09\x52\xc6\x97\x69\x03\xdd\x1a\xaf\xde\x0e\x7d\xc7\x6e\xe8\xcc\x19\xfe\xba\xc7\x43\xbb\x23\x20\xb0\x65\x7b\xb8\xdd\xf6\x90\x4f\x03\x56\x8c\xd2\xc1\xf3\xe4\x52\x2e\xd9\x66\xe0\x56\x44\x2c\xce\xf9\x1c\x93\xf4\x85\x07\x2b\xb5\xf7\x45\x9d\x0f\xac\x28\x1b\x31\x45\x12\x8f\x08\x26\xa7\xe4\x72\x44\x11\x04\x18\x03\x49\x45\x13\xad\x41\xab\x8b\x08\x17\xd2\x15\x37\x8d\x34\x0b\xeb\xaa\x02\xeb\xf0\xe6\x05\xdb\xf5\x07\x64\x19\x47\x68\xb7\x98\xef\x85\x46\x27\x06\x55\x6d\xa5\xfd\x75\x3c\x1e\xf0\xff\x4d\x60\xd2\x09\xd2\x11\xed\x03\xa9\x05\xe8\x82\x51\x4b\x88\x85\x5b\xae\x1d\xb7\xda\x21\xe7\x80\x28\xb2\x58\x8c\xd7\xf5\xb1\x9b\x78\x09\xcf\xa4\x6a\xb8\x6d\xb8\xbb\x00\xaf\xfe\x38\xd2\xc2\x6d\xbf\x6e\x53\x0d\x7f\x04\x88\xbe\x4e\x7f\x18\x59\xbc\x8f\x90\xa7\x81\xb2\x00\x1a\x14\x16\x61\x82\x43\xe3\x69\x8b\x67\xcb\x45\x50\xc5\x52\x12\x91\xdc\x7a\x92\x97\xb1\x29\x9c\x0f\x2a\x66\x69\x3d\xfd\xa5\x5c\xc5\xbd\x37\x46\x3a\x45\x79\x2b\x1b\x4b\xf9\x4a\xe6\x18\x93\x53\x98\xa3\xb7\x8b\x81\xfa\xee\xd0\xde\x81\xb2\x8c\xeb\x46\x53\xc2\xe2\x0b\x83\xfa\x01\x0f\x6b\x64\xf1\x00\x5d\xda\xb0\xe7\xfb\x43\x3f\xf6\xcf\x4f\xc3\x12\x12\x64\x99\x0d\xf2\x88\x59\x50\x5a\x4f\x0a\xb8\x13\x31\x45\xab\x75\x94\x59\x40\x02\x4e\xa1\x2f\x3e\xfa\xad\xd3\x5a\x07\x20\x75\x0c\xa4\x0e\x23\xb4\x0e\xb2\x48\xb4\x0e\x8c\x77\x56\x7b\xa1\x73\x94\xda\x65\x13\x92\x3e\x1f\x69\xd6\xc9\x9b\xc0\x7d\xe4\x5e\x95\x84\x91\x93\x01\x09\x26\x29\x9b\xa2\x8e\xda\xf0\x94\xa2\xf0\xa8\xd2\x72\xb1\xf2\x88\xec\xf5\x71\xcd\x44\x60\x52\x36\x4a\x36\xd2\x2d\x9c\xdc\xfe\xde\xb5\x63\x4b\xd9\xeb\xcb\x1a\x3b\x29\xe9\xcc\x4d\x6b\x97\x70\x09\x9c\x84\xd1\xf0\xa5\x5d\xc2\x37\x47\x12\x05\xb0\x61\x3c\xee\xdb\xcc\xc6\x4d\xdb\xad\x87\xc9\xd4\x8d\x40\x51\xac\xda\x85\xd3\x94\x32\x7b\x2f\x9e\xbd\xdb\xcf\xd4\xe2\xd7\xdd\xed\x3f\xa6\x91\xb9\xba\xd9\xb5\x6d\x34\x5f\x98\xd1\xb7\x93\x81\x63\x1b\x18\x36\x8c\x3c\x5c\xdb\x11\x9c\x2e\x21\xbb\xff\xcc\x69\x56\xa0\xc8\x64\x05\x02\x2a\x57\x4a\x8c\x1a\x22\x51\xa2\x94\xa0\xb5\x11\xde\x14\xc0\x10\x93\x50\xd1\x3f\x21\xc3\xc4\xa4\x0c\xd8\x12\xb0\x24\x8e\xc9\x04\x2f\xc8\x79\x3b\xaf\x40\x41\xe6\x02\x7c\xb6\xda\x26\xf7\x99\x85\xf9\xea\x80\xa4\x60\x82\x7c\x9b\x6a\xb4\xb9\x80\xd7\x1e\x30\xf5\x5d\xae\xc2\xa9\xa5\xd0\x9d\x4b\xfb\x13\x49\xe9\x98\x36\xc7\xfd\x70\xe7\x0e\x36\x70\x2a\x2d\x13\xb3\xf2\x04\x48\x57\x29\x95\x4c\xb8\xc8\xa9\x99\xea\x86\x73\x82\xe8\x14\x99\xe4\xa4\x93\x3b\x9f\xd6\xec\xfd\xb0\xbc\x9c\xb8\x6b\x8c\x5f\x3e\xe4\x55\x58\x0f\x4a\xa0\x7b\x41\x4d\x25\xd2\x69\x7c\x66\xd7\x95\x67\xf3\x99\x6d\x1f\x2e\xa0\x26\x29\x1a\x0a\x79\x7c\xe9\xe4\x3e\xd5\x3f\x69\x8b\xd0\x55\xe1\x00\x89\x2a\xe2\x58\x7d\xf6\x80\x44\x03\xc6\xf6\x54\x4a\xb3\x23\xd2\xd9\x94\x54\x11\x32\xe4\xfa\x07\x82\x4c\x5c\x07\x45\xb5\xaa\x54\x29\x52\x2a\x94\x02\x9c\x4d\x92\x30\xc9\xcf\xca\xec\xaf\x46\xe9\x07\x4a\x0c\x13\xa5\x9a\x8c\x6a\xb9\xca\xc3\xc2\x94\x7e\x19\x90\x55\x47\x4d\x71\x2d\x6d\x61\x98\x5c\xc3\xa6\xcd\xc8\x70\x8b\xbb\x09\xc6\x69\x4a\xa2\xd2\x87\x2a\x91\x22\x9c\xf5\xa0\xbd\xe1\x11\x1c\x4a\xfa\x27\x9b\xe8\x45\x4e\x05\x13\x50\xb6\x17\x4f\xe8\x01\x45\x4d\xac\x95\x72\x64\xe8\x4c\xf6\x46\x29\xa5\xc9\xa7\x10\x7e\x8a\xa4\x29\x28\x5b\xb2\x0a\xf4\x6c\xd1\xca\xa7\x32\x22\x77\x27\x7c\xa3\x5d\x23\xf5\xc2\x8c\x7e\xde\x1c\xfa\x1b\x8a\x6f\x6c\xd7\xa7\x77\xaf\xa6\xca\x87\x75\x7d\xae\xf9\x69\x8b\x27\x02\x9a\xed\xbb\xce\x4f\x45\x73\xbb\x14\x17\xe8\xf9\xf0\xd0\x18\x4a\x47\x4b\xf7\x7c\x7e\x3a\x7e\xe8\xec\x1c\xbb\x4c\xab\x9f\x93\x85\x18\xa1\xed\xf0\x30\x3c\x67\xef\xf0\xb6\xa2\x3b\xab\xe1\x13\x6c\x8c\x19\xc0\xd0\xa5\xa3\xb4\xa4\x62\x2c\xda\x5b\x59\x5c\x94\x25\x59\x15\x79\xa4\xac\x41\x89\xff\x09\x0d\xba\xd9\x9d\x8a\x0b\xf4\x11\x84\x68\x4c\x68\x64\x58\x18\xd6\x77\xed\xb6\x9a\x05\x7a\x94\x9e\x8d\x10\xb7\x53\xcb\x9b\x91\xac\x12\x0e\xc3\xa9\x7a\x52\x44\x23\xa3\x93\x39\x93\x8a\x3c\x05\x23\xe7\x11\xbd\x0a\x29\xd9\x20\x54\x26\xff\xe5\x74\x74\xe1\x09\x88\x42\x00\xf0\xbc\x40\x22\xec\xdc\x93\x1f\x01\x52\x13\xa9\x55\x7a\x11\xad\x45\x4e\xfe\x31\x18\x29\xca\x25\x92\x12\x21\x12\xaa\xe1\x8b\x07\x3c\xc8\xec\x53\xfb\x70\xa2\xdb\x96\xd2\x72\x4f\x5b\x9a\x72\x90\x84\x9c\xe9\x4a\x69\x2a\x58\x23\x19\xee\x40\x8a\xa2\xbc\xae\x45\x10\x68\x9f\x10\xe2\xb8\x42\x32\x05\xa0\xb4\x30\x46\x1b\x99\xa2\x01\x02\x91\x7c\x70\x45\x3a\x6d\x01\xab\x7c\xf4\x82\x21\xee\x93\xda\x48\x53\x7c\x5b\x7e\x0e\xf1\xb7\x16\xba\x91\x1d\x90\xf6\xe1\x38\x75\x48\xd9\x4d\x9d\x13\x95\x43\xbf\x23\x25\xd5\x8f\x61\x66\xdf\x5f\xff\xf8\x8f\x29\x10\x5e\x9d\x46\x8c\x32\x8f\x13\x5e\x51\xe4\xa1\x10\x04\x51\x29\x4f\xd7\x96\x6b\x49\x2e\x2f\x19\xa8\x1d\xe9\x84\x3e\xf1\xf2\x04\x62\x73\x2b\xb0\x0b\x11\xb3\x8d\x90\x17\x28\x90\x62\x7f\x43\x9f\x9c\x66\x46\x35\xc6\xad\x0f\xb0\xdf\x6c\xdb\xcf\x9a\x2a\x5f\xd0\xd3\xdc\xdd\x7e\x36\x9d\xd9\xfc\x2f\xe1\xbb\xe5\x34\x7e\x5d\xba\x1b\x74\x7d\x83\x38\xbe\x60\xeb\x7e\x64\xed\x38\x69\xa9\xed\xd6\x0c\x28\x35\x0d\xfb\x2d\xbd\x72\xdc\x9f\xa8\xa1\xd9\xb5\xcb\xf9\x1a\x9a\x1b\x57\x2f\xe1\xbc\x43\x53\xbd\xdd\xe2\x3d\xbb\x9a\x7f\x88\x05\x7b\x3b\x4c\x66\xfb\x05\x45\xc0\x5d\xff\xbe\xa2\xaa\x3d\x6b\xc2\x41\x51\x2f\xb6\xe3\x70\x22\x2c\x13\xc8\x01\x73\xf2\x73\x5e\xd2\x85\x2b\x99\xae\xa7\x28\x06\x09\x99\x87\xa2\x82\x26\x84\x46\x69\xfe\x04\x58\x92\x1b\x1d\x78\x76\x60\x33\xa7\x92\x48\x50\xa9\xec\xa9\x3a\xa6\xe3\xc4\x18\x4d\x12\x92\xbc\x84\x35\x0b\x4e\x5e\x1f\x4e\xf2\x8c\x6c\xc4\xd2\x03\xa2\xd7\x69\x3c\x92\x80\x6e\xd9\x0e\xde\x21\x1b\x37\x78\x3f\x74\x9d\x46\x79\x55\x62\xa7\x41\x9a\x3d\x0c\x3d\xbf\x96\x9d\x5b\x2d\x5f\xc2\xda\x85\x46\x71\xaa\x96\x96\x8e\x7b\x53\x74\xab\xd5\xeb\x40\xbe\xa1\x8e\x28\xee\xba\x41\x78\x57\x40\x6d\xdb\xe1\xd4\x96\xc3\x6c\x7f\x7c\xbe\x03\x9f\x5b\x86\x2d\x89\x69\xd2\x52\xa8\x03\x25\x61\xe8\x71\x81\xa9\xf8\x87\xb6\x50\x8d\x7b\xec\x0f\x64\x1f\x3e\xe4\xaa\x88\xf4\x1c\xff\x78\xa2\x96\xe6\x86\xa3\x27\xf4\xee\x66\x16\x62\x97\x98\x88\xfb\x46\xde\x7b\xf2\x05\x21\xd5\xde\xdd\x16\xbb\xf5\xb8\x79\x19\x61\xa0\x4c\x34\x6c\xfa\xe3\x36\x7f\xdb\xef\xf6\x7d\x47\x30\x7e\xd9\xd7\x93\x63\xe3\xa1\x4d\xef\x26\xff\xf7\x5f\x4e\xff\x0b\x2d\xd7\xb9\xb7\x19\x9c\x8d\x6d\x76\x26\xbc\x90\xa9\xf0\x0d\xbf\xc4\x98\x29\x2b\xb6\xed\xd3\x94\xad\x2a\x98\xbe\x94\x6d\xdb\xd5\xdb\x85\xde\x9f\xd8\x6f\x9d\xdb\x0d\xfd\x5f\x80\xa6\x67\x6f\xee\x8e\xf9\xe1\x8e\xc9\x97\x3b\x1c\x61\xce\x5d\x93\xe7\xde\xb1\x30\xc2\x7a\xba\x2b\x73\xd5\x76\xed\xb8\x7a\xf3\xf8\x5d\x1c\xc6\xeb\x3a\x7e\xab\x87\xd8\xc3\x40\x1b\x7b\x75\x81\x9c\xfa\xf1\xa4\xc4\x15\xbf\xe2\x5f\xfb\xac\xce\xad\x3f\x1f\x9c\x95\x58\x5d\x22\x20\x3e\xfc\xa6\xd8\x65\x72\xbb\xa7\x7c\xd5\x73\xef\xf9\x7b\x70\xdc\x32\xdd\x35\xb7\x7a\xbc\x4f\x9f\xfd\xfe\xec\x3f\x01\x00\x00\xff\xff\xda\xac\x91\x10\xdd\x2b\x00\x00")

func _3_json_bytes() ([]byte, error) {
	return bindata_read(
		__3_json,
		"3.json",
	)
}

func _3_json() (*asset, error) {
	bytes, err := _3_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "3.json", size: 11229, mode: os.FileMode(420), modTime: time.Unix(1431985517, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var __4_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x00\x02\xa5\xd4\xbc\xb2\xcc\xa2\xfc\xbc\xdc\xd4\xbc\x92\x62\x25\x2b\x85\x68\xb0\x28\x08\x54\xc3\x59\x60\x75\x25\x95\x05\xa9\x40\x79\x25\xa0\xea\x92\xd2\xc4\x1c\x25\x1d\x54\xe9\xc4\x94\x94\xa2\xd4\x62\x90\x09\xa8\xfa\xc0\x92\x19\xf9\xc5\x25\x79\x89\xb9\x60\xfd\x29\xa9\x65\x06\x86\x4a\x28\x6a\x6a\xd1\xcc\x2a\x28\xca\x2f\xcb\x4c\x49\x2d\x02\x5b\x97\x98\x5e\x94\x98\x57\x82\xd0\x80\xa4\x98\x86\x2e\x34\xa2\xc4\x85\x60\x56\x2c\x57\x2d\x17\x20\x00\x00\xff\xff\xdc\x1e\x51\xf8\x63\x01\x00\x00")

func _4_json_bytes() ([]byte, error) {
	return bindata_read(
		__4_json,
		"4.json",
	)
}

func _4_json() (*asset, error) {
	bytes, err := _4_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "4.json", size: 355, mode: os.FileMode(420), modTime: time.Unix(1431118616, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var __5_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x94\xc1\x6a\x84\x30\x10\x86\xef\x79\x8a\x25\xe7\xae\x69\xa5\xf4\xd0\x57\x29\x3d\xc4\xcd\xb0\x09\x68\x92\xce\xcc\x5a\x96\xc5\x77\x6f\x8c\x20\xc4\x4a\x6f\x45\x0f\xce\x49\xe7\xff\xc7\xdf\xf9\x08\x79\x88\x53\x2a\xd9\x86\xab\xbb\x9c\x89\x35\x03\xc9\xf7\xd3\x47\xee\x8e\xf5\x98\x9f\xb2\x8f\xef\x11\x92\x2e\x1b\xe7\x35\xde\xe5\x53\xa9\x3a\x93\xb4\x72\x22\xf7\x7b\x40\x72\xc1\x8f\x83\x75\x55\x57\xaf\xb2\xb0\x0c\x8b\xaf\x44\xcd\x76\xb4\xaa\x1b\xa1\xa2\x94\xa4\x2c\x73\x34\xcb\x30\xf0\xbd\xc3\xe0\x3b\xf0\xbc\x9e\xaa\x8d\x41\x20\x5a\x15\xb3\xc1\x06\x62\xaf\xbb\xbc\x50\xc4\x60\xce\xdf\xd0\x3c\xbf\xc8\x5f\xe6\xe1\xcf\xdf\x6d\x5d\x83\x1a\xdd\x82\xdb\x2c\x2b\xe0\xcb\xb4\x80\xea\x82\xb9\xb5\x40\x2a\x4d\x44\x1b\x2b\x0a\x65\xd4\xa7\x58\x89\xf8\x0f\xfe\x6f\x3b\xe4\x9f\x8e\xde\x15\x0e\xf4\x5b\xa0\xff\xd2\x07\xf7\x2d\xb8\x1b\xe8\x8f\xdb\x66\x3b\xf4\xf5\x1e\xd0\x8b\xe9\x7d\x10\x3f\x01\x00\x00\xff\xff\x66\xed\x8c\x2b\x88\x07\x00\x00")

func _5_json_bytes() ([]byte, error) {
	return bindata_read(
		__5_json,
		"5.json",
	)
}

func _5_json() (*asset, error) {
	bytes, err := _5_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "5.json", size: 1928, mode: os.FileMode(420), modTime: time.Unix(1431118616, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var __6_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x95\x51\x4e\x83\x40\x10\x86\xdf\x39\x05\xd9\x67\x91\x61\xd5\x44\x39\x47\x9f\x34\x8d\x19\xdc\x49\x4b\x52\x16\xca\xac\x24\xa6\xe1\x50\x5e\xc4\x33\x09\x5b\x4b\x42\x59\xc1\x47\x69\x99\xf0\x00\xbb\x1f\x3f\x33\x3f\x3f\xe1\xe0\xf9\x4d\x09\x85\x06\x99\x0c\x8b\xd8\x7f\xb1\x2b\x6d\x1d\xba\x33\xcb\x68\xcc\xa8\xd9\x17\x61\x85\x65\xb8\x4b\x93\x30\xfb\xe0\xfd\x4e\xdc\xf4\x29\xd2\x55\x5a\xe6\x3a\x23\x6d\x1a\xb8\x2f\x61\x01\x54\xaa\x24\x66\xe7\xa6\x05\xb6\x39\x9b\xd3\xb3\x8a\x32\x57\x81\x4a\x20\x12\x03\xb6\xee\xad\xd4\x67\x6d\x14\x68\xb6\xd3\xcd\xf2\x7b\x32\x98\xda\x3d\xfd\xc0\x05\x2c\x8a\x57\x95\x9c\x09\x76\xd4\x86\x34\x71\xda\x0a\x8b\xaf\xcf\xa9\xde\xd7\x9e\x63\x8a\x7f\x60\x3e\x1b\xdc\xd0\xcc\x8c\x7f\x2b\x09\x0d\x05\x26\x3d\xc2\x12\xa2\x87\x00\xa2\x00\x9e\x56\x20\x63\x88\x62\x09\xb7\x00\xf0\x3c\xfd\xde\xdc\x5d\x58\x68\xca\xc4\x0e\xfc\x63\x92\x4f\x55\xbb\x9b\xb2\x52\x3f\x1f\x68\x6b\xd8\xd0\xdb\xce\x14\x7f\x3d\x22\xc1\x1a\x8b\x69\x63\x9c\xf7\xd7\x33\x4d\xf0\x1e\x2f\x24\xbe\x8f\x2b\x79\x17\xdf\x37\x87\x5c\xe2\x3b\x6a\xcc\x45\xc5\x57\x51\x35\xbb\x3f\xdf\x92\xe0\x5f\x25\xae\x34\xc1\x72\x49\xf0\x00\x5a\x12\x3c\x92\x60\xef\x78\x5d\x7b\xdf\x01\x00\x00\xff\xff\x15\x03\x09\x4c\x9d\x0c\x00\x00")

func _6_json_bytes() ([]byte, error) {
	return bindata_read(
		__6_json,
		"6.json",
	)
}

func _6_json() (*asset, error) {
	bytes, err := _6_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "6.json", size: 3229, mode: os.FileMode(420), modTime: time.Unix(1431118616, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var __7_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x93\xb1\x6e\x84\x30\x0c\x86\xf7\x3c\x05\xca\x5c\x48\x69\xc5\xd2\x57\xa9\x3a\x04\x62\x91\x48\x90\xa4\xb1\x4b\x85\x10\xef\x7e\x21\xc3\xe9\xe0\x18\x60\x3e\x3c\x19\xff\xbf\x71\xfc\x29\x99\x58\x16\x83\x77\xae\x35\x4d\x8e\x24\x09\x90\x7f\x65\xdf\xa9\xba\xc4\x74\xcf\x92\x8f\x46\x0f\x51\xe7\x9d\xa9\x83\x0c\x23\x7f\x5b\xcb\x46\x45\x71\xdd\x92\xea\x03\x04\x34\xce\x2e\x9d\x55\x51\x15\xe5\x27\x5f\x79\xe6\xcd\x6f\xbc\x24\xbd\x78\x05\x50\x23\x34\x91\x57\xa2\x77\xea\xaf\x03\x14\x71\xb0\xd7\xbe\x40\xb7\x1d\x0d\x76\x30\xc1\xd9\x1e\x2c\xed\x9f\x41\x2a\x15\x00\x71\x57\x4c\x06\xed\x90\xac\xec\xd3\x7e\x3e\x38\x95\xff\x43\xfd\x5e\xf2\x27\xf3\xcc\xf6\xbf\x1e\xb6\x78\x51\x6a\xf1\xfe\xb4\x70\x01\x3b\x0e\xec\x57\x5e\xb4\x8e\xd3\x52\x30\x5c\xef\xf1\x2c\xb0\x8f\x13\xc0\x52\xf6\xc3\x66\x76\x0b\x00\x00\xff\xff\xf3\x26\x00\x02\x16\x06\x00\x00")

func _7_json_bytes() ([]byte, error) {
	return bindata_read(
		__7_json,
		"7.json",
	)
}

func _7_json() (*asset, error) {
	bytes, err := _7_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "7.json", size: 1558, mode: os.FileMode(420), modTime: time.Unix(1431118616, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var __8_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x97\xdd\x8e\xb2\x30\x10\x86\xcf\xb9\x0a\xd2\x63\xfd\x28\x3f\x82\x7a\x2b\x5f\x3c\xa8\xb4\x51\xb2\xd8\xd6\xb6\x88\xc6\x78\xef\x5b\x20\xab\xcb\xdf\x92\x6c\x82\xd9\x06\x7a\x60\x84\x79\xdb\x99\xe9\x3c\x79\xa3\x77\xcb\xd6\x0b\x70\xc1\x62\x22\x25\x91\x60\x6b\xff\x2f\x5f\x15\xeb\xfe\xfc\x56\x8a\x52\x76\x48\xe2\xa5\x54\x48\x55\x3a\x1b\x38\x17\x24\x9c\x3c\xcf\x1d\xc4\x39\x58\xe8\xe7\x4c\x0a\x47\xee\x13\xea\x1c\x95\xe2\x18\xd8\xbb\x45\xfd\x88\x4c\x12\xa1\xb7\x02\xbd\x07\x34\x42\x07\xc1\x32\xde\x13\x8b\x73\x5c\x44\xca\x74\x22\xa3\x55\xba\xba\x84\x27\x85\x24\xf4\x5c\xaf\x11\x20\xf4\x92\x08\x46\x4f\x84\x2a\x2d\xa8\xb7\x54\x0a\x10\xc6\x42\xf7\xde\x19\x2c\x05\x47\x26\x15\x45\x27\x52\x94\xa0\x2f\x0a\x2f\x73\xb2\x87\x2e\x68\x89\x1f\xb5\x37\x8f\x46\x1d\x69\x22\x15\xa1\xb5\xfb\xfd\x5a\x3d\x79\xd5\x8d\x57\x39\x99\x50\x8d\x76\x5f\x6d\x0b\xa6\x58\x35\x0c\x15\xf3\xd6\x85\xbf\x74\xc5\x19\x5b\x7b\x0d\x07\xaa\xde\x59\x1d\xf5\x0f\x73\x50\xcc\x3d\x4d\xf6\xe4\x4a\x62\xe7\x74\x93\xe7\xf4\x87\xd9\x97\xf1\xde\xe9\x77\x46\x31\x52\x48\x12\xf5\x64\x40\xa7\x72\x3a\x85\x15\x05\x5e\x08\xdf\x81\x01\x36\x95\x02\xdf\x87\x61\xbb\xee\xf6\x9e\x81\x8a\x32\x9a\x5c\x7b\x2b\x42\xea\xd8\x1e\x57\xf5\xf9\x4f\xb2\xf8\x63\xe8\xe2\x7e\x09\xe2\x5f\x33\x24\x3f\x88\xc2\x60\x54\x14\x75\xfb\x07\x62\x24\x86\xd3\x30\xa3\x68\x15\x46\x33\x00\x0d\xdd\xec\x43\xef\xf6\xa1\x75\xe0\xaf\x46\xc5\xf0\x8c\x8c\x64\x70\x1a\x26\xe4\x06\xeb\x8d\x3b\x8f\xff\xbb\x6e\xb6\xa0\x37\x5b\x90\xb7\xda\x44\x70\x54\x06\x31\xb9\x18\xfa\x8b\x7c\x1a\x2e\xe4\x43\xe8\x8f\xfb\xbf\xcc\x5c\x02\x66\x23\x7a\x1e\x31\xba\x11\x41\xcf\x0d\x47\xc7\xd0\x33\x12\xc3\x89\x18\x51\xe4\x6f\x66\x02\x9a\x3a\x23\x8d\xc8\xaa\x9e\x1f\xd6\x67\x00\x00\x00\xff\xff\x77\x55\x92\x6c\xca\x16\x00\x00")

func _8_json_bytes() ([]byte, error) {
	return bindata_read(
		__8_json,
		"8.json",
	)
}

func _8_json() (*asset, error) {
	bytes, err := _8_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "8.json", size: 5834, mode: os.FileMode(420), modTime: time.Unix(1431118616, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"1.json": _1_json,
	"2.json": _2_json,
	"3.json": _3_json,
	"4.json": _4_json,
	"5.json": _5_json,
	"6.json": _6_json,
	"7.json": _7_json,
	"8.json": _8_json,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"1.json": &_bintree_t{_1_json, map[string]*_bintree_t{
	}},
	"2.json": &_bintree_t{_2_json, map[string]*_bintree_t{
	}},
	"3.json": &_bintree_t{_3_json, map[string]*_bintree_t{
	}},
	"4.json": &_bintree_t{_4_json, map[string]*_bintree_t{
	}},
	"5.json": &_bintree_t{_5_json, map[string]*_bintree_t{
	}},
	"6.json": &_bintree_t{_6_json, map[string]*_bintree_t{
	}},
	"7.json": &_bintree_t{_7_json, map[string]*_bintree_t{
	}},
	"8.json": &_bintree_t{_8_json, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


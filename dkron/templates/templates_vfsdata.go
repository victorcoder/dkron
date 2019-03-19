// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates contains project templates.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 1, 30, 20, 19, 45, 952386193, time.UTC),
		},
		"/dashboard.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.html.tmpl",
			modTime:          time.Date(2019, 1, 30, 20, 19, 45, 949543295, time.UTC),
			uncompressedSize: 4100,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xb8\x12\x03\xda\x00\x93\xf4\xb0\xb7\x4e\x36\x90\x25\x05\xd6\x75\x4d\x83\x36\x18\x30\x0c\x43\x70\x12\xcf\x12\x63\x8a\xe4\x48\xca\x8d\x61\xf8\xbb\x0f\x92\x2c\x59\x72\xe4\x24\xf6\xbc\x27\x91\xc7\xbb\xdf\xfd\x25\xef\x14\xbf\xb9\xfe\x72\x75\xf7\xe7\xed\x07\xc8\x7d\x21\x67\x93\xb8\xfa\x80\xca\x02\x34\x66\xca\xf8\xc2\x6a\xc5\x66\x93\x49\x9c\x13\xf2\xd9\x04\x20\xf6\xc2\x4b\x9a\xad\xd7\x10\x5e\xe9\xa2\xd0\x2a\xbc\xc1\x82\x60\xb3\x89\xa3\xe6\xa4\xe2\x71\xa9\x15\xc6\x83\xb3\xe9\x94\xf5\x38\x2f\x9d\x23\xef\x6e\xd1\xe7\xb0\xd9\x44\x0f\x2e\x42\x95\x95\x12\x6d\x58\x08\x15\x3e\x38\x36\x8b\xa3\x46\xf2\x14\x90\xc0\xea\xd2\xd3\x08\xd4\x31\x58\x1c\x5d\x9e\x68\xb4\xfc\xa9\x39\x52\xa8\x05\x58\x92\x53\xe6\xfc\x4a\x92\xcb\x89\x3c\x83\xdc\xd2\xfc\x19\xc8\xd4\xb9\x28\x93\x2b\x93\x8b\x54\xab\xc0\x19\xa1\x14\xd9\x30\x75\x8e\x1d\x67\xd7\xc3\x3f\x25\xd9\xd5\x01\xe7\x4e\xb6\x2c\xd1\xda\x3b\x6f\xd1\xd4\xc0\x8d\x55\x47\x18\x35\x14\x3f\xa3\x5d\x68\xcc\x09\x31\xe2\x3f\x8d\xd7\xd1\x49\x66\x58\x91\x2e\x5c\x8e\xdf\xbb\xc5\x09\x21\x1a\xc7\x38\xad\xce\xbb\x22\x7f\x82\x79\xae\xb8\x6b\x4e\x85\xb0\x56\xdb\x48\x8a\xa4\xb7\x3d\xd2\xe9\xc3\x38\xa7\x39\xde\xc3\x2b\x34\xa7\xe8\x01\x97\xd8\x08\xf6\x96\xff\x1d\x1b\x39\xd7\x2a\xe2\xc2\x19\x89\xab\x08\x4b\xaf\x2d\xcd\x2d\xb9\xfc\xf4\x77\xa9\x14\xc1\x41\xef\xf7\xd2\x94\x6b\xeb\xd3\xd2\x43\xf5\x4a\xbc\x98\xa9\x39\x2e\x2b\xbe\x50\xa4\x9a\x81\x5f\x19\x9a\x32\x51\x60\x46\xd1\x63\x50\xcb\xef\x15\xfd\x99\x30\x77\x9e\x57\xf8\x00\x4b\xb4\x70\xfd\xe9\xeb\x97\x9b\xfb\xcb\xdb\x8f\xf7\xb7\x97\x77\xbf\xc2\x14\x06\x0a\x6e\x3f\x6e\xd1\xd9\xcf\xb5\xc4\x0f\xef\xe6\xa5\x4a\xbd\xd0\x0a\xde\x5d\xc0\xba\xa6\x55\xd4\xb7\x7f\x71\xf4\x18\x78\x9d\x65\x92\xa6\xcc\x6b\x2d\xbd\x30\xec\xef\xb7\x17\xe1\x76\xfd\xee\xa2\x66\xde\x54\x9f\x5d\x0c\xe3\xa8\xe9\x49\x93\x38\xd1\x7c\x55\x7b\xcd\xc5\x12\x52\x89\xce\x4d\x99\xc2\x65\x82\x16\x9a\x4f\xc0\x69\x8e\xa5\xf4\xed\x76\x2e\x1e\x89\x07\x5e\x1b\x06\x56\x57\x4a\x15\x2e\x45\x86\x95\x6d\xac\x71\xaf\x0f\x95\x6a\xe5\x51\x28\xb2\xdb\x33\x80\xf8\x4d\x10\xc0\x95\x96\x12\x8d\x23\x0e\x3b\x69\x08\x82\x8e\xe7\x89\x31\x41\x65\x6f\x0f\x65\x8b\xf3\xe1\xd1\xa0\xe2\x64\x21\x29\xbd\x1f\x40\x00\xc4\x5b\x5a\x93\x92\x66\xc3\xf6\x50\x9b\xc0\x31\x18\x44\x31\xdd\x1a\xd7\x92\xd1\x66\xe4\xa7\x2c\xdc\xca\x74\xc7\x3b\x55\x55\x7e\x0d\xaa\x16\xdc\xd9\x40\x2b\xb9\x62\xb3\xbb\x1a\xb1\xe7\x63\x1c\x55\x7c\x07\x05\xeb\x4e\x97\xa0\xad\xcb\xfd\x7f\x66\x8c\xa3\x26\x24\x75\x79\xf6\x42\xfa\x19\x85\x82\x7a\x1c\x19\x46\xb3\x97\x92\xc4\xa2\xe2\x43\xf7\xb1\x3d\x93\x3a\xd3\x23\x57\xa6\x2d\xe7\x59\x2c\x8a\xec\x85\xfb\x2f\x8a\x2c\xaa\x87\xa7\x20\x41\x9e\x51\x90\x7c\x0f\x8d\xca\x18\x44\x7d\x8d\x00\x63\x63\x14\xf6\xfd\xe3\x62\xd9\x15\x54\xb3\x99\x3c\x2d\x9d\xd7\x57\x60\x9b\x77\x18\x29\x80\x5d\xe4\x46\xd1\x00\xe2\x52\xb6\x70\x05\x0a\x15\x54\x6f\x8c\xab\x98\xdb\x6b\xa5\x70\x39\x8c\xa8\x14\x43\x77\x63\x7c\x26\xaa\x03\x4e\x80\x58\xb4\xca\xba\x01\x0a\x76\xa3\x54\x37\xa8\x55\x45\x21\xf6\x65\xaf\xdb\xd3\xa1\xf6\x7e\x68\xab\xed\xd0\xbc\x63\xac\x8d\x1e\x74\xe2\x8e\x34\xd9\xa3\x5b\xb8\x51\x73\x7f\xd3\x89\x7b\xbd\xa5\x71\x54\xca\xbd\xac\x7d\x23\xb4\x69\xfe\x23\xdc\xec\x12\x87\x8a\xc3\x57\x32\x1a\x9a\x2c\x1d\xc8\xe3\x30\x79\xed\xd2\x8a\x2c\xf7\xaf\xcc\x64\xee\xbd\x79\x1f\x35\xb5\x1e\x8a\xaa\x7b\x6c\x9f\x9a\xfb\x44\xa2\x5a\x1c\x19\xa3\x44\xeb\xc5\x78\x46\x75\x7a\x62\x88\x7a\x77\xa8\x5b\x76\x8b\xf5\xda\x53\x61\x24\x7a\x82\xfa\x99\x27\xe5\x19\x84\x9b\xcd\xb3\xcd\xa4\xe9\x1e\x89\xf6\x5e\x17\x2d\x4d\xa8\x25\xd9\xee\x3a\x8d\xdc\xbc\x7a\xf2\x67\x63\x57\xb3\x6b\x2f\xe0\xe9\xd1\x07\x29\x29\x3f\x6c\x12\x83\x74\xbd\x90\x16\xb3\xdf\x1c\xe8\xd1\xb3\x91\xbf\xb5\xfe\xcb\xf3\x07\x59\x57\x95\x4c\xf5\xf8\x98\xe3\x6e\xc8\xb8\xba\xcf\x54\x24\x64\xdf\xf7\x75\x34\xa4\xee\x8d\x3b\x8b\x9a\xdf\xeb\x76\x3a\x50\xd3\x90\xce\xab\xe6\x9b\xd7\x16\x33\x82\x04\xd3\x05\x29\x3e\xd0\xf7\x4b\x43\x3b\x9f\xb2\x4f\xb4\x72\x06\x53\x1a\x68\x69\x89\x2f\xa9\x79\x65\xd9\xc7\x51\x33\x30\x55\x13\x54\xfd\xcf\xff\x6f\x00\x00\x00\xff\xff\x9f\xc8\x9f\xd2\x04\x10\x00\x00"),
		},
		"/executions.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "executions.html.tmpl",
			modTime:          time.Date(2019, 1, 30, 20, 19, 45, 950491732, time.UTC),
			uncompressedSize: 2615,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x51\x6f\xdb\xb8\x0f\x7f\xef\xa7\x60\xf5\x0f\x1a\x07\xff\xda\x5e\x87\x3e\x6d\xb6\x81\xdd\x70\x77\x40\xb1\xdb\x01\xd7\xc3\xbd\x1e\x64\x8b\xb1\xb4\x3a\x92\x21\xd1\x4b\x03\xcf\xdf\xfd\x20\x39\x4e\x9c\x26\xc5\xd0\x3e\xb8\x12\x49\x91\x3f\xfe\x48\x51\xe9\x7b\x81\x6b\xa5\x11\x58\x65\x34\xa1\x26\x36\x0c\x57\x99\x50\xdf\xa1\x6a\xb8\x73\x79\x10\x73\xa5\xd1\xc2\x64\x00\xba\x8e\xfd\xda\x9a\xa6\x41\x9b\xb3\x5f\x9f\xb1\xea\x48\x19\xed\x3e\x93\x6d\x58\x71\x05\x30\x77\x60\xcd\x36\xc8\xf6\x52\x25\x72\xb6\x41\xe7\x78\x8d\xac\xc8\x52\xa1\xbe\x87\x03\x87\xc5\xe5\x93\xf2\xae\x38\x86\xc9\x52\x79\x37\xc9\xdf\x17\x8f\xd2\x6c\x95\xae\xe1\x0b\x77\x04\x77\xef\xde\xc1\x83\x29\x01\x0f\xc6\xb0\x36\x16\xbe\x99\x12\xfa\x3e\x79\x30\xe5\x57\xbe\xc1\x61\x08\x87\x01\x32\x0e\xd2\xe2\x3a\x67\xff\x1b\x93\x6a\x54\xf5\x94\x33\xdb\xe9\x07\x53\x46\xcb\xf9\x81\xe5\x2a\x58\x38\x69\xb6\x39\xbb\xb6\x9d\xd6\x4a\xd7\xff\x7e\x33\x25\x03\x52\xd4\x60\xce\xfe\xea\x34\x2b\x32\x35\x61\xaf\x9b\x5d\x2b\x55\x65\x34\x1c\x56\x71\xdb\xf0\x1d\xdb\x87\xf6\x7f\xdc\x2a\x1e\x4b\x25\x04\xea\x9c\x91\xed\x02\x21\xaa\xc8\x52\x5e\x9c\x03\xfc\x89\xef\xca\xd4\xa0\x8d\xdd\xf0\x26\xb6\xaa\x96\x14\xbb\x56\x69\x8d\x76\x06\x7b\x8e\xfa\x34\x50\x96\xca\xf7\xb3\x2a\x5c\x01\xf4\xbd\xe5\xba\x46\x58\xd4\x4f\xf0\x21\x87\xe4\x97\xdd\xef\xd6\x74\x6d\x60\x2e\x93\xf7\xc7\x62\x7c\x80\xbe\x5f\xd4\x4f\xc3\x90\xa5\xf2\x3e\xf8\x20\x5e\x36\x38\x41\x1d\x37\xe1\x1b\x3b\xb2\xaa\x45\x31\x95\x94\xec\x21\x4b\x92\xc5\x83\x29\xb3\x94\xe4\x5c\xf4\x48\xdc\x12\x0a\xf8\x44\x2f\x35\xbf\x29\xad\x9c\xbc\xa8\xfa\x6a\x04\xbe\x94\xfd\xd9\x51\xdb\x9d\x59\x3e\x76\x55\x85\xce\x1d\xc5\x59\x3a\x41\x3a\x64\x8f\x4f\xb7\xb0\xc0\x67\x4f\x41\xa4\xb4\xc0\x67\x58\x24\x81\x08\xe7\x99\x59\xed\x1b\xe9\x24\x15\x51\xf4\xfd\x02\x9f\x8f\xad\x93\xa5\x24\xce\xd5\xfb\xe4\x3e\xd1\x6b\x06\x53\x8e\xaf\x5b\xf8\x54\x2f\x87\x38\xf4\xd8\xb1\x7f\xfa\x1e\x66\xa8\x60\x18\x62\x2f\xa9\x9f\xa6\x15\x86\xd5\xc6\x08\xde\x30\x10\x9c\x78\x4c\xa6\xae\x7d\x63\x9f\xc8\xb8\xad\x91\xde\xe8\xae\x98\xb5\xfc\x88\x7c\x2c\x08\xfc\x00\x32\x8f\x64\xfd\xe5\xfd\x01\x64\x3b\x5d\x71\xc2\x61\x48\x92\x64\x7e\x05\x2e\xd3\x37\x16\x6f\x9e\xfa\x58\xbe\x71\x79\x1d\xc7\xf0\x87\x0f\x0e\x71\x3c\x9b\x3d\xfb\xa6\x0c\xb0\x60\xcd\x05\xb2\x30\x8e\xde\xc4\x0d\xf1\x32\x74\x42\xce\xe2\x3b\x06\xd6\x78\x8a\x84\xe2\x8d\xa9\xa7\xab\x1d\xae\x75\xc3\x4b\x6c\x1a\x14\xe5\x2e\x67\x9b\x5d\xc0\xf2\xc5\x8b\x0e\x6c\x9c\x01\x8a\xf7\x5e\x26\x9f\xa6\xea\x36\x7e\xdc\xce\xaa\x79\x76\x64\x1a\xc9\x73\x8a\xcf\xad\x24\x72\x81\xf6\xc4\x08\x20\x2b\x3b\x22\xa3\x81\x76\x2d\xe6\x6c\xdc\xb0\xc3\xd0\x6f\x8c\xc3\x7d\xcd\x85\x72\x1b\x75\x70\xc6\x66\xe9\xe5\xec\x73\xb0\x2b\x32\xd7\x72\x7d\x69\x9c\xdd\x90\xda\xa0\xfb\x98\xa5\xde\xa0\xc8\xd2\x31\xcc\x0b\x20\xf2\xfe\x14\x6e\x18\xa8\x63\x65\x4e\x99\xfb\x47\xe1\x16\xcc\xd8\x3b\x7e\xaa\x9f\x5d\xb4\x71\x02\x1d\x1c\x4f\x6f\xca\xeb\xcc\x94\x46\xec\x5e\xf2\xd2\x5a\x2c\x5e\x69\x53\x1f\xc3\xab\xdf\x16\x64\x6d\x0c\xbd\x8d\xfe\x92\x34\x94\xa4\x63\x81\x6b\xde\x35\x74\xb9\x10\x45\x20\xff\x12\xa7\x2f\x30\x9d\x6c\x67\x9b\xd9\xb2\xef\x51\x8b\x71\xb8\xa7\x61\x58\x17\x57\x47\xe1\xde\x2e\x73\x95\x55\x2d\x79\xcd\x22\x9a\xba\x73\x95\x58\xe4\x62\x17\xad\x3b\x5d\xf9\xe7\x00\xa2\x15\xf4\xc1\xa5\x5a\x43\xb4\x55\x5a\x98\x6d\xd2\x98\x8a\x7b\x65\x22\xb9\x93\x70\x73\x03\x8b\x8b\x9a\x55\xd2\xa0\xae\x49\x4e\x1e\xe0\x55\xbb\x40\x40\xb4\xf4\xcf\xda\x72\xf5\x31\x18\x8f\x93\x78\x11\x2d\x47\xe5\x72\x95\x18\x3d\x5a\x24\xa5\xdb\xcb\x6e\xe1\x08\x13\x8f\x51\x3c\x52\x5f\x06\xb3\x86\x08\x13\x8b\x0d\x27\x14\x7f\x87\x49\xb7\x82\xeb\x1c\x58\xa7\xc7\xdf\x48\x82\x1d\x0f\x01\x5c\x4c\x2e\x87\xc5\x99\x8f\x84\x13\xd9\x68\xe9\x07\xf1\x84\x76\xc2\x3b\xec\xf7\x73\x42\x3d\xee\xf1\x16\xfd\x14\xb9\x54\x8e\x8c\xdd\x25\x6d\xe7\xe4\x23\x71\xc2\x88\xb1\x5b\x98\x5c\x25\xe1\x22\xdd\x9e\x01\x6d\x39\x49\xed\x07\xdd\xff\xcf\x54\x0e\xb9\xad\xe4\x44\x69\xf8\xef\xbf\x59\x3a\xd5\x7e\xea\x89\xff\x02\x00\x00\xff\xff\x83\x7f\x99\xd3\x37\x0a\x00\x00"),
		},
		"/index.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.html.tmpl",
			modTime:          time.Date(2019, 1, 30, 20, 19, 45, 950890894, time.UTC),
			uncompressedSize: 1090,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x8f\xda\x30\x10\xbd\xf3\x2b\x46\xbe\xd0\xaa\x85\x48\xbb\xb7\x95\x13\xa9\xed\xa9\x52\xb5\xaa\xd4\xfe\x01\xc3\x0c\xc4\xc2\xd8\xd1\x78\x60\x8b\xd2\xfc\xf7\xca\x21\x21\xa6\xb0\xab\xbd\xc4\xcf\xef\x3d\x8f\x3d\x1f\x69\x5b\xa4\x8d\xf5\x04\x6a\x1d\xbc\x90\x17\xd5\x75\x33\x8d\xf6\x08\x6b\x67\x62\x2c\x7b\xda\x58\x4f\x0c\xa3\x01\xfc\x76\x91\x30\x07\xe7\x88\x4b\xf5\xdd\x23\xfd\xf9\x26\xec\x54\x35\x03\x68\x5b\xa1\x7d\xe3\x8c\x10\xa8\x28\x46\x0e\x51\xc1\xb2\xeb\x66\x00\x79\x54\x0e\x2f\xbd\xfb\x9a\x6d\x8c\x27\x07\xfd\x77\x81\xb4\x31\x07\x27\x83\xeb\x8e\x6f\x51\x93\x41\xeb\xb7\x17\x07\x80\xae\x1f\xaf\x2d\x62\xc5\x91\xaa\x7e\xd8\x23\xe9\xa2\x7e\xbc\xc4\x2a\xd0\x1e\x5f\x0f\xbc\x0a\x78\xca\xa3\xb2\x5d\xef\x62\x6d\x5e\x60\x04\x8b\xd0\x88\x0d\x3e\x96\x6a\x00\x6a\x92\x36\x64\xe4\xc0\x14\x4b\x35\xa2\x4c\x8c\xc4\x36\x49\xe7\x55\x55\xba\x18\xa5\x3b\x4f\xbb\xc0\x01\xbc\x5e\xc2\xfa\xa1\x7a\x0e\x48\x51\x17\xf5\xc3\x40\x89\x59\x39\x1a\xad\xe7\x4d\xff\x5d\x44\x61\xdb\x10\x4e\x75\x15\xce\x52\x95\xba\x7a\x36\x7b\xd2\x85\xd4\xd7\xec\x17\x44\xa6\x18\x6f\x85\x9f\x81\xe5\x96\xfd\xd5\x77\xfe\x96\xff\x6d\xb6\x57\xac\x2e\xa6\xeb\xb5\x70\x1a\x2d\xa6\x86\x8c\x94\x6a\x4f\xfb\x15\x31\x58\x0f\x67\x14\xe1\x2f\x04\x46\xe2\xaf\xa7\xa7\x79\x7a\xe4\x3c\xef\x91\x60\x3a\xbb\xb2\x1e\xc7\x93\xcb\xe4\x49\x25\x16\x7c\xdb\x97\x52\x7b\x8f\x2f\x65\xfa\x1e\xdf\x39\xf7\x3b\xce\x69\x33\x74\x32\xcb\xf6\xc3\x8e\x4e\x9f\xe1\x68\xdc\x81\x3e\x4e\x39\x2f\x53\xbd\x54\x7e\x0e\x40\xc7\xc6\xf8\xe9\xd2\x1d\x9d\xe0\x13\xcc\x9f\x60\x9e\x6e\x4c\x5a\xf5\x9f\xa3\x0f\x7a\x11\xf3\x37\xe4\xbf\x41\xdf\x0b\xbc\xed\x8b\x2e\xfa\xb9\xc9\xe6\x70\x58\xda\x96\x3c\x76\xdd\xec\x5f\x00\x00\x00\xff\xff\x65\x81\xbc\x4b\x42\x04\x00\x00"),
		},
		"/jobs.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "jobs.html.tmpl",
			modTime:          time.Date(2019, 3, 12, 16, 50, 23, 955540252, time.UTC),
			uncompressedSize: 7191,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xef\x8e\xdb\xb8\x11\xff\xbe\x4f\x31\x61\x16\xb1\x8d\x44\xf6\x05\x38\xa0\xc0\x46\xf2\xe1\x90\xb4\x48\x17\xbb\x97\x43\xb3\xed\x87\x16\x45\x8e\x12\xc7\x16\x53\x9a\x14\x48\x6a\x77\x0d\xc7\x4f\xd4\x47\xe8\xb7\x7b\xb2\x82\xa4\x64\xfd\xb1\xfc\x27\x69\x82\xa2\xc5\xed\x07\x4b\x22\x87\xa3\x99\xdf\xfc\x66\x38\xd4\x6e\x36\x0c\x17\x5c\x22\x90\x4c\x49\x8b\xd2\x92\xed\xf6\x22\x66\xfc\x1e\x32\x41\x8d\x49\xfc\x30\xe5\x12\x35\xd4\x02\x20\x97\x91\xbb\xd7\x4a\x08\xd4\x09\xb9\x56\xe9\x0d\x37\xf6\xb5\xd5\x82\xcc\x2f\x00\x36\x1b\x8b\xab\x42\x50\x8b\x40\x8c\xa5\xb6\x34\x04\xa6\xdb\xed\x05\x40\x5b\xaf\x56\x0f\x5e\x1a\x20\x5e\x28\xbd\x0a\xb7\x5d\x11\x37\x1e\x2d\xb5\x2a\x0b\x52\x4f\x77\x05\xb8\x2c\x4a\xbb\x27\x71\x50\x26\xa2\x8c\x29\x49\xe6\x31\xaf\x27\x97\x62\x5d\xe4\x3c\x53\x12\x76\x77\x91\x41\xaa\xb3\x9c\xcc\xe3\x19\x9f\xc7\x33\xc6\xef\x3b\x9a\xbd\x3a\xb0\xeb\x02\x13\x62\xf1\xd1\x92\x8e\xb1\x15\x2c\x04\x0a\x41\x33\xcc\x95\x60\x0e\xa0\x3f\x70\x61\x51\x7b\xdc\x56\x8a\xa1\x48\x48\x78\xc7\xb5\x4a\x3f\xc7\xec\x96\x24\x40\x4c\x21\xd7\xb8\x48\xc8\x53\x02\x8c\x5a\x1a\x59\xb5\x5c\x0a\x4c\xc8\x4a\x31\x2a\xea\x31\xaa\x97\x68\x13\xf2\xf4\xa3\x4a\xef\xaa\x98\xdc\x86\x79\xcb\xad\x93\xfe\x09\x1f\x1c\x1e\x1d\xd5\xee\xef\x08\x3e\x85\x70\x01\xa5\x9a\xd3\x28\xe7\x8c\xa1\x4c\x88\xd5\x25\xd6\x80\xd1\x8e\x47\x5d\xfc\x3a\x8f\xad\x87\x78\xd6\x50\x20\xb6\x34\x15\x58\x1b\x10\x1e\xfc\x6f\x64\xac\xe6\x05\xb2\x1d\x10\xb1\xcd\x91\xb2\x96\x76\xab\x3b\xef\xb6\xf9\xfc\x5a\xa5\xf1\xcc\xe6\xfd\xe1\xf7\x59\x8e\xac\x14\x38\x38\x57\x66\x19\x1a\x33\x34\xf5\x7b\xad\x95\x1e\x9a\xb8\xa1\xc6\x02\x3e\x62\x56\x5a\xae\xe4\x90\xc4\x4f\xf8\x78\x42\xe2\xbd\xcf\x94\xa1\x99\x1f\x33\xb7\xa6\x37\x15\xcf\x1a\x6f\xdd\x4c\x0b\x89\xd8\x2e\x94\xb2\x6d\x5c\x18\x64\x4a\x98\x82\xca\x84\xfc\xee\x20\xe3\x8a\x52\x88\x48\xf3\x65\x6e\xfb\x54\x2b\xc5\x4e\x86\x2e\xb9\xa4\xce\x9c\x9e\x0c\x40\x2c\xb8\x2f\x0c\x41\x70\xc3\xb8\x71\x41\x63\x57\x90\x95\x5a\xa3\xb4\x3f\xd3\x25\x42\x92\xc0\x77\xdb\xbd\xa5\x3b\x36\x07\x05\x3c\xfb\x47\x42\x0a\x8d\xf7\x6e\xc9\x78\x42\xe6\xbf\xfe\x13\x7e\xd6\x78\xdf\xe3\x56\xf0\x5c\xf0\xf9\xc5\xb0\x25\x1a\x0b\xa4\x36\x21\x12\xb8\x04\x4d\xe5\x12\xc7\x05\x5d\x22\xfb\xa3\xc5\x95\x99\x0a\x94\x4b\x9b\xbf\x68\x5b\xd7\x79\x80\xe7\xb0\xa4\xc5\x04\xc8\x9e\xad\x2d\x27\x69\x66\xf9\x3d\x5e\x81\x74\x8e\xb5\x16\x6f\x49\xcb\x13\x83\xb6\x76\xe4\x98\xdf\x29\x97\xcc\x19\xfb\x1c\x5e\x92\x7e\x1a\x9d\x72\x75\x1f\xf4\x71\xcb\x9a\x89\xb3\x6e\xcf\x75\x88\xe0\xe5\x79\xa1\x90\xf8\xb8\xf3\xc0\xf3\xf8\xd7\x7f\x1d\xb4\xaf\x3b\x52\x8a\xa3\xc5\xc0\xb2\x16\x83\xdb\x9c\x8d\x6d\xaa\xd8\xba\x93\xd9\xed\x88\x7e\x54\xa9\x8b\x69\xe3\xd2\xdf\x5a\xde\xfe\x1d\x3e\xc1\xc2\x17\xdc\xab\x03\x55\xd6\xb2\x79\x4c\x9d\xbe\x50\x40\x3f\xaa\xd4\xcc\x36\x9b\x5f\x36\x1b\xf8\xa8\xd2\xa9\xa4\x2b\x84\xed\xf6\x97\xed\x76\xb6\xcb\x57\x43\x9a\x08\xd5\x32\x21\x48\x6d\x27\x2a\xe5\x5d\x51\x53\x95\x1a\x72\x86\x68\xa8\x3c\x1f\x32\x55\x4a\x7b\x86\x3c\xba\x72\x74\x44\xba\x13\x8c\xdd\xca\xb1\x5b\x2a\xa8\xb1\x1f\xaa\xf7\xc1\x1c\x76\x43\x5e\x25\xfc\x00\x7b\x32\x57\x3d\x99\x09\x7c\x72\x5b\x0c\x5e\x8d\xd6\xeb\xf5\x3a\xba\xbd\x8d\x18\x83\xb7\x6f\xaf\x56\xab\x2b\x63\xe0\xaf\xa3\x2e\xe2\xa7\x3c\x71\x0c\x3b\xa9\x70\x40\x4b\x8f\x6f\xed\x5c\x70\x1c\xf9\x04\xa1\xfd\x78\xed\x86\xc8\xde\x9e\x16\x76\xac\xe3\x76\xee\x6d\xba\x3b\xda\x3c\x0d\x94\xb9\xe4\x92\xe1\x63\x20\x4c\xd4\xd9\x7b\x8f\xed\xc7\x07\xd7\xf6\x92\xaa\xda\xa5\xff\xc2\xf1\x61\xbf\xe0\x1e\xed\x62\x70\x8d\x91\x2a\x50\x1e\xdc\xa9\x7b\xa9\xda\xcb\xe7\x67\x32\x35\xc5\xab\x83\x1d\x47\x53\x1c\x74\x29\xaf\x55\x3a\xae\xb3\x62\xe2\xe7\x4c\xae\x1e\x12\xf2\x44\x97\x52\x72\xb9\xfc\xb0\xef\xed\xae\xff\xf8\x53\x29\x3f\xbf\xff\xa0\xeb\x73\xfb\x8f\x8e\xd5\x27\xfa\xbe\x4c\x2d\x41\x2a\xbd\xa2\xd5\x2e\x18\x99\x82\x4b\x89\x7a\x70\x07\x08\x1e\x1e\x71\x70\xd8\x9a\xb3\x61\x0d\xec\xe9\x21\x5b\x61\x76\xe7\xe7\x4e\xb9\x53\xd0\xd2\xe0\xbe\xed\xe7\xe2\x76\xc2\xd2\x2f\xa7\x78\xc4\x50\xa0\xdd\xb3\xac\x61\x8d\x9f\x3f\x41\x9b\x37\x41\xc7\x09\x08\xac\xa6\x26\xff\x72\x08\x76\xfe\x7e\x03\xe6\x1c\x73\x72\xb8\x91\x6e\xd7\xa2\x5e\xf3\xd7\x6c\x96\xf1\xcc\xb7\xca\x55\x9f\x10\x3f\x89\x22\xf0\xfd\x3e\x04\xc0\x20\x8a\x2a\x39\xd7\xf7\xfd\x47\x3b\x6a\x8d\x87\x8f\x29\x2c\x28\xdb\x45\x94\xb3\x84\x9c\x0a\xbe\xeb\xe8\xfd\x64\x42\xa2\x97\x04\xb4\x72\x41\x65\x9c\x0a\xb5\xac\x52\x5b\xd0\x14\x85\x40\x96\xae\x13\xb2\x5a\x7b\x27\x6e\xdc\x10\x19\x3a\x28\x56\xba\xab\xf5\x95\x36\x95\x95\x2b\x77\x5c\x1d\x3e\x3a\x86\x25\xf5\x91\xf6\x50\x4f\x1c\xa4\x1c\xbe\xbd\xf2\xfb\xa3\x46\x58\xab\x12\x4c\xa9\xf1\x87\xc3\x3d\xce\x90\x36\xd7\xe8\xa0\xee\xf7\xd8\x69\x69\xad\x92\xb5\x64\x6a\x25\xa4\x56\x46\x0c\x17\xb4\x14\xb6\x5d\x19\x02\x82\xbd\xca\x70\x4f\x45\xd9\xa4\x45\x8f\x76\x3e\x2f\x19\x37\x2b\xbe\x33\x82\xcc\x5f\x2b\xb9\xe0\x7a\x15\xcf\xc2\x8b\x87\xad\x09\xc7\xdc\xf0\x40\x0e\xda\x36\xac\x5f\x28\x83\x43\xda\xcf\x3e\x12\xfa\xdb\x3e\x8f\xbf\x29\x81\x8f\x52\xb7\x46\xf5\x7f\x8a\xb9\xee\x58\x78\x90\x6b\x83\xd1\xcd\x5c\xdc\x86\x63\xda\x72\x2f\x21\x3e\xbe\x64\x1e\xbb\x53\xe5\x39\x15\xf6\x99\xe5\x2b\x34\xaf\xe2\x99\x5b\x30\x3f\x40\xbb\xfc\xfb\xae\xf9\xbe\xda\x87\xb0\x74\x90\x6c\xba\xc7\x91\x6b\x8d\x60\x04\xcf\xa1\xdd\x97\xe7\xdf\x1f\x61\xdc\xc1\x0c\x77\x5a\xb9\xe4\x81\x4d\xd7\x46\xc9\xc4\x2a\x77\x71\x89\xd6\x3f\xbc\xc5\x16\x1f\x2d\xd5\x48\xa1\xe4\x51\xa6\x18\xae\xb8\x6f\x9d\x3b\x4f\x91\x2a\xac\x49\x08\x32\x6e\x95\x7e\x57\x34\x27\x89\xea\x2b\x50\xf5\x1a\xdf\xda\x56\xea\x3e\xd3\xee\xa3\xb5\xe4\x5b\x64\xef\xe7\x54\xaa\xb2\x70\x0d\x7d\x55\xa9\x9c\xa3\x4d\xa1\x0a\x53\x67\x15\xaa\x3f\x7b\xd1\xaf\x58\x49\xae\x55\x0a\xf5\xc7\xb0\xa1\xaa\x72\xa0\x2e\xec\x7f\x42\xfb\x2d\xf7\xff\xcb\xb9\xff\x5a\x23\xb5\x38\xfa\xf2\x8c\xff\x36\x39\x5d\xd3\xe4\xff\x39\xb7\x33\x0f\x7d\x95\xdb\x6d\x87\x9b\x1c\x0f\xd1\x39\xaf\x19\xf1\xa2\x5f\x25\xc7\xab\xbb\xfa\x62\x32\xcd\x0b\xff\x5d\xe9\x72\x5c\xe7\xd6\x64\xaa\x91\xb2\xf5\x78\x51\x4a\xff\x65\x15\xc6\x13\xd8\x78\x35\xb3\x19\xdc\x70\x63\x51\xc2\x42\x69\xb0\x39\xba\x5d\x05\xb4\x7a\x30\x60\x15\xd0\xa2\x40\xaa\xbd\xe0\xe5\x78\x34\xf5\x9d\xf6\x68\x32\x75\x64\x1c\x8f\xde\xbc\xbb\x7d\x5f\xa6\x56\xa3\xab\x0e\x7c\xc1\x91\x8d\x5e\x40\xf3\x06\xac\x5f\x01\xc0\x17\x30\xc6\x69\x38\x27\x4d\xfd\x39\xe1\xed\xdd\xed\x4d\xfd\x55\x6e\x0e\xdf\x35\xa2\x41\xf8\x81\x4b\xa6\x1e\xa6\x42\x65\xfe\xcb\xeb\x34\xa7\x26\x87\x67\xcf\xe0\x72\x70\x66\x52\x69\x6a\x6b\x81\x83\xb2\x3e\x00\xe3\x91\x3b\x94\x8c\x26\xcd\x71\x6f\x7b\xd1\xbe\x6e\x27\xaf\x2e\x6a\x7c\xee\x72\x6e\xc0\xe4\xaa\x14\x0c\x52\x74\x3d\x97\x83\xc9\xaa\x02\x04\xde\xa3\x80\x1a\x64\x07\x98\x43\xc6\x5d\x9d\xc4\xa2\xb4\xa5\x46\x78\xf3\xee\x16\x50\xe0\xca\xd4\xfa\x6c\x4e\x2d\x50\x8d\x20\x95\x05\x2a\x7c\x64\xa0\xd0\x68\x9c\x0e\x25\x81\xf8\x11\x32\xad\x60\x6f\x62\xa8\x64\x30\x7b\x9a\x9a\xe0\xc5\x31\xbc\x5d\xba\x28\x8f\xbb\x46\xc7\x54\x76\xe7\xe1\x9f\xc0\x93\x04\x48\x29\xc3\xbf\xc0\x18\x69\x63\x36\x88\x7a\x02\x97\x7b\x3a\xa6\xd4\x5a\x3d\x1e\xb9\xa3\x62\x03\x61\x0f\xb8\xbe\xe1\xa1\x3e\x9e\x34\x3d\xe7\xc6\x2a\xbd\x9e\x16\xa5\xc9\xdf\x5b\x6a\x71\x4c\xc8\x8b\x1d\xc4\x53\x5f\x22\x5f\xec\x59\x5a\x50\x9b\xfb\xaf\x98\xcf\xf7\xa6\x42\xfb\x5b\x59\xb9\xf5\x57\xf7\x1b\xcf\xea\x3c\xd9\x6c\x50\xb2\xed\xf6\xe2\xdf\x01\x00\x00\xff\xff\x5e\xa7\x4b\x69\x17\x1c\x00\x00"),
		},
		"/status.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "status.html.tmpl",
			modTime:          time.Date(2019, 1, 30, 20, 19, 45, 953447095, time.UTC),
			uncompressedSize: 1053,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x93\x31\x4f\xfb\x30\x10\xc5\xf7\x7c\x8a\x93\x77\xff\x23\xfd\x67\xd7\x1b\x0c\xac\xb0\x57\x8e\x7d\x09\x46\x57\xbb\xf2\xd9\xa5\x28\xca\x77\x47\x4d\x9d\x52\x24\x50\x2b\x44\x59\x72\xd1\xbb\x77\x97\x5f\x9e\xe5\x71\x74\xd8\xfb\x80\x20\x38\x9b\x5c\x58\x4c\x53\xa3\x9c\xdf\x81\x25\xc3\xbc\x12\x29\xbe\x0a\xdd\x00\xcc\x9a\x77\x2b\xb1\x41\x66\x33\xa0\xd0\xaa\x75\x7e\xa7\x9b\xa5\x7c\x37\x53\xb5\xad\x09\x48\x30\x3f\xa5\xc3\xde\x14\xca\xb3\xe7\x0b\x97\xec\xa2\x7b\x83\x23\x8e\x9c\x95\xea\x04\x50\x85\x16\x2b\x79\xce\xd2\x07\x3a\xb0\x57\xf4\xc5\x05\xa0\xc8\x2f\xbe\x8c\xfb\x2c\x2d\x86\x8c\x09\x6c\x24\xb9\x71\xf2\xff\x99\x13\x40\x71\x4e\x31\x0c\x10\x06\xd9\xf9\xe0\x56\xe2\x25\x76\xfc\x8f\x30\x0c\xf9\xf9\xf0\x97\xc7\xf6\xa7\x89\xad\x56\xbc\x31\x44\xfa\x29\x66\x43\xf0\x10\x3b\x56\xed\x51\xf9\x40\x68\xc9\xff\x12\x10\x17\x6b\x91\xb9\x2f\xb4\x3e\xb0\x5d\x80\x5a\xbe\x53\x03\xac\xc3\x42\x3f\x9e\xb6\xdc\x1c\xb8\x37\x9e\xd0\xfd\x00\xd6\x99\x30\x60\x12\xfa\x7e\x5e\xf0\x57\xc1\xae\x6d\x2c\x21\x5f\x3a\xeb\xb3\xfc\xee\xf6\x68\x4b\xf6\x31\xdc\x94\x0e\x53\x8a\xe9\x3a\xb6\x9a\xd7\x95\x5c\xaa\x2d\xb5\x53\x2f\xef\xe9\xa5\x96\x71\xc4\xe0\xa6\xa9\x79\x0f\x00\x00\xff\xff\x33\xc9\xe6\xe6\x1d\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/dashboard.html.tmpl"].(os.FileInfo),
		fs["/executions.html.tmpl"].(os.FileInfo),
		fs["/index.html.tmpl"].(os.FileInfo),
		fs["/jobs.html.tmpl"].(os.FileInfo),
		fs["/status.html.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}

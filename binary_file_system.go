package main

import (
	"errors"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
	"fmt"

	"github.com/c2h5oh/datasize"
)

var (
	errDirListingUnsupported = errors.New("not supported")
	kindNameRegexp           = regexp.MustCompile("/files/([^/]+)/([^/]+)")
)

type binaryFileInfo struct {
	fileSize int64
}

func (bfi *binaryFileInfo) Name() string {
	return ""
}

func (bfi *binaryFileInfo) Size() int64 {
	return bfi.fileSize
}

func (bfi *binaryFileInfo) Mode() os.FileMode {
	return os.FileMode(os.ModePerm)
}

func (bfi *binaryFileInfo) ModTime() time.Time {
	return time.Now()
}

func (bfi *binaryFileInfo) IsDir() bool {
	return false
}

func (bfi *binaryFileInfo) Sys() interface{} {
	return nil
}

type binaryFile struct {
	fileType string
	fileSize int64
}

func (bf *binaryFile) Close() error {
	return nil
}

func (bf *binaryFile) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		if bf.fileType == "sparse" {
			p[i] = 0
		} else if bf.fileType == "random" {
			p[i] = byte(rand.Int())
		} else {
			fileTypeBytes := []byte(bf.fileType)
			fmt.Println(fileTypeBytes)
			var j int
			for j = 0; j < len(fileTypeBytes) && i+j < len(p); j++ {
				p[i+j] = fileTypeBytes[j]
			}
			i += j
		}
	}
	return len(p), nil
}

func (bf *binaryFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, errDirListingUnsupported
}

func (bf *binaryFile) Stat() (os.FileInfo, error) {
	return &binaryFileInfo{fileSize: bf.fileSize}, nil
}

func (bf *binaryFile) Seek(offset int64, whence int) (int64, error) {
	return -1, nil
}

type binaryFileSystemImpl struct{}

func (bfsi *binaryFileSystemImpl) Open(name string) (http.File, error) {
	m := kindNameRegexp.FindStringSubmatch(name)
	fileType := m[1]
	fileSize := m[2]

	var fileSizeValue datasize.ByteSize
	fileSizeValue.UnmarshalText([]byte(fileSize))
	return &binaryFile{fileType, int64(fileSizeValue.Bytes())}, nil
}

func createBinaryFileSystem() http.FileSystem {
	return &binaryFileSystemImpl{}
}

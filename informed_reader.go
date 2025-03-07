package memory_streams

import (
	"io"
	"io/fs"
)

type InformedReader struct {
	io.Reader
	info fs.FileInfo
}

func NewInformedReader(reader io.Reader, info fs.FileInfo) *InformedReader {
	return &InformedReader{
		Reader: reader,
		info:   info,
	}
}

func (reader InformedReader) Info() fs.FileInfo {
	return reader.info
}

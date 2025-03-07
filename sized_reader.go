package memory_streams

import "io"

type SizedReader struct {
	io.ReadCloser
	totalSize int64
}

func NewSizedReader(reader io.ReadCloser, totalSize int64) *SizedReader {
	return &SizedReader{
		ReadCloser: reader,
		totalSize:  totalSize,
	}
}

func (reader *SizedReader) Size() int64 {
	return reader.totalSize
}

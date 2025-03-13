package streams

import (
	"io"
)

// SizedReader associates it with a fixed total size
type SizedReader struct {
	reader io.ReadCloser
	size   int64
}

// NewSizedReader creates a new SizedReader with the given io.ReadCloser and total size
func NewSizedReader(reader io.ReadCloser, totalSize int64) *SizedReader {
	if reader == nil {
		panic("parameter reader cannot be nil")
	}

	return &SizedReader{reader: reader, size: totalSize}
}

// Size returns the total size in bytes of the data expected to be read
func (sr *SizedReader) Size() int64 {
	return sr.size
}

// Read reads data into the provided byte slice
func (sr *SizedReader) Read(p []byte) (int, error) {
	return sr.reader.Read(p)
}

// Close closes the reader
func (sr *SizedReader) Close() error {
	return sr.reader.Close()
}

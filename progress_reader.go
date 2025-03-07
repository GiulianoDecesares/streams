package memory_streams

import "io"

type ProgressReader struct {
	reader    io.ReadCloser
	progress  chan float64
	total     int64
	read      int64
	notifying bool
}

func NewProgressReader(reader io.ReadCloser, total int64) *ProgressReader {
	progressReader := &ProgressReader{
		reader:    reader,
		progress:  make(chan float64),
		total:     total,
		read:      0,
		notifying: true,
	}

	return progressReader
}

func (reader *ProgressReader) Read(bytes []byte) (amount int, err error) {
	amount, err = reader.reader.Read(bytes)
	reader.read += int64(amount)

	if reader.notifying {
		if reader.read >= reader.total { // Total could be 0, or reader.read == reader.total
			reader.progress <- 1.0 // Total reached instantlly

			reader.notifying = false
			close(reader.progress)
		} else { // Notify partial progress normally
			progress := float64(reader.read) / float64(reader.total)
			reader.progress <- progress
		}
	}

	return amount, err
}

func (reader *ProgressReader) Close() error {
	if reader.notifying {
		close(reader.progress)
		reader.notifying = false
	}

	return reader.reader.Close()
}

func (reader *ProgressReader) Progress() <-chan float64 {
	return reader.progress
}

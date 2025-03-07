package memory_streams

import (
	"io"
	"time"
)

type ReadNotification func(read int64, bytesPerSecond float64)

type NotifierReader struct {
	reader       io.Reader
	read         int64
	notification ReadNotification
	startTime    time.Time
}

func NewNotifierReader(reader io.Reader, onProgress ReadNotification) *NotifierReader {
	return &NotifierReader{
		reader:       reader,
		read:         0,
		notification: onProgress,
		startTime:    time.Now(),
	}
}

func (reader *NotifierReader) Read(bytes []byte) (amount int, err error) {
	amount, err = reader.reader.Read(bytes)
	reader.read += int64(amount)

	if reader.notification != nil {
		elapsed := time.Since(reader.startTime).Seconds()
		bytesPerSec := float64(reader.read) / elapsed // Speed in bytes per second
		reader.notification(reader.read, bytesPerSec)
	}

	return amount, err
}

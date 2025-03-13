package streams

import (
	"io"
	"sync"
	"time"

	"github.com/GiulianoDecesares/bandwidth"
)

// ReadNotification is a callback function type that provides progress updates
type ReadNotification func(read int64, speed bandwidth.Bandwidth)

// NotifierReader wraps an io.ReadCloser and notifies progress updates
type NotifierReader struct {
	reader       io.ReadCloser
	read         int64
	notification ReadNotification
	startTime    time.Time
	lastNotify   time.Time
	mu           sync.Mutex
}

// NewNotifierReader creates a new NotifierReader
func NewNotifierReader(reader io.ReadCloser, onProgress ReadNotification) *NotifierReader {
	now := time.Now()

	return &NotifierReader{
		reader:       reader,
		notification: onProgress,
		startTime:    now,
		lastNotify:   now,
	}
}

// Read reads data into the provided byte slice and sends progress updates
func (reader *NotifierReader) Read(data []byte) (int, error) {
	amount, err := reader.reader.Read(data)

	reader.mu.Lock()
	defer reader.mu.Unlock()

	if amount > 0 {
		reader.read += int64(amount)
	}

	if reader.notification != nil {
		// Prevent division by zero
		elapsed := max(time.Since(reader.startTime).Seconds(), 0.01)
		bytesPerSec := float64(reader.read) / elapsed

		// Reduce notification frequency (e.g., notify every 100ms)
		if time.Since(reader.lastNotify) > 100*time.Millisecond {
			reader.notification(reader.read, bandwidth.New(bytesPerSec, bandwidth.BytePerSecond))
			reader.lastNotify = time.Now()
		}
	}

	return amount, err
}

// Close closes the underlying reader
func (reader *NotifierReader) Close() error {
	if reader.reader != nil {
		return reader.reader.Close()
	}

	return nil
}

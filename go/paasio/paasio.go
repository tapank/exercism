package paasio

import (
	"io"
	"sync"
)

type Counter struct {
	rbytes int64
	wbytes int64
	rcalls int
	wcalls int
	r      io.Reader
	w      io.Writer
	mu     sync.Mutex
}

type rCounter Counter
type wCounter Counter
type rwCounter Counter

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &wCounter{w: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &rCounter{r: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &rwCounter{r: readwriter, w: readwriter}
}

// Read implements ReadCounter
func (rc *rCounter) Read(p []byte) (n int, err error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	n, err = rc.r.Read(p)
	rc.rbytes += int64(n)
	rc.rcalls++
	return
}

// ReadCount implements ReadCounter
func (rc *rCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.rbytes, rc.rcalls
}

// Write implements WriteCounter
func (wc *wCounter) Write(p []byte) (n int, err error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	n, err = wc.w.Write(p)
	wc.wbytes += int64(n)
	wc.wcalls++
	return
}

// WriteCount implements WriteCounter
func (wc *wCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.wbytes, wc.wcalls
}

// Read implements ReadWriteCounter
func (rw *rwCounter) Read(p []byte) (n int, err error) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	n, err = rw.r.Read(p)
	rw.rbytes += int64(n)
	rw.rcalls++
	return
}

// ReadCount implements ReadWriteCounter
func (rw *rwCounter) ReadCount() (int64, int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	return rw.rbytes, rw.rcalls
}

// Write implements ReadWriteCounter
func (rw *rwCounter) Write(p []byte) (n int, err error) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	n, err = rw.w.Write(p)
	rw.wbytes += int64(n)
	rw.wcalls++
	return
}

// WriteCount implements ReadWriteCounter
func (rw *rwCounter) WriteCount() (int64, int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	return rw.wbytes, rw.wcalls
}

package circular

import "errors"

// Circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.
type Buffer struct {
	data       []byte
	size, len  int
	rpos, wpos int
}

func NewBuffer(size int) *Buffer {
	if size < 1 {
		return nil
	}
	return &Buffer{data: make([]byte, size), size: size}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.len == 0 {
		return 0, errors.New("buffer empty")
	}
	b.len--
	v := b.data[b.rpos]
	if b.rpos++; b.rpos >= b.size {
		b.rpos = 0
	}
	return v, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.len == b.size {
		return errors.New("buffer full")
	}
	b.len++
	b.data[b.wpos] = c
	if b.wpos++; b.wpos >= b.size {
		b.wpos = 0
	}
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.len == b.size {
		if b.rpos++; b.rpos == b.size {
			b.rpos = 0
		}
		b.len--
	}
	b.len++
	b.data[b.wpos] = c
	if b.wpos++; b.wpos >= b.size {
		b.wpos = 0
	}
}

func (b *Buffer) Reset() {
	b.len = 0
	b.rpos, b.wpos = 0, 0
}

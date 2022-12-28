package ylog

import (
"bytes"
"sync"
)

var defaultBufferPool *BufferPool

type BufferPool struct {
	pool *sync.Pool
}

func (s *BufferPool) Get() *bytes.Buffer {
	return s.pool.Get().(*bytes.Buffer)
}

func (s *BufferPool) Put(buf *bytes.Buffer) {
	s.pool.Put(buf)
}

func SetBufferPool(bp *BufferPool) {
	defaultBufferPool = bp
}

func init() {
	SetBufferPool(&BufferPool{
		pool: &sync.Pool{New: func() any {
			return new(bytes.Buffer)
		}},
	})
}
package ylog

import (
"sync"
)
type EntryPool struct {
	pool *sync.Pool
}

func (s *EntryPool) Get() *Entry {
	return s.pool.Get().(*Entry)
}

func (s *EntryPool) Put(e *Entry) {
	s.pool.Put(e)
}

func newEnP() *EntryPool {
	return &EntryPool{
		pool: &sync.Pool{
			New: func() any {
				return &Entry{}
			},
		},
	}
}

package fifo

import (
	"unsafe"
)

type FifoItem uint64

type Fifo struct {
	head int
	tail int
	data []FifoItem
	size int
}

func New(size int) *Fifo {
	s := new(Fifo)
	s.data = make([]FifoItem, size+1, size+1)
	s.size = size
	s.head = 0
	s.tail = 0
	return s
}

// Allocate specified number of bytes from the FIFO starting at tail
func (f *Fifo) Allocate(size int) (address unsafe.Pointer, ok bool) {
	// if I fail to allocate
	if tail > head {
		if b.tail + size < f.size {
			b.tail += size
			return  
		}
	} else {
		
	}
}

func (f *Fifo) Free(address unsafe.Pointer, size int) ok bool {
}

func (f *Fifo) Len() int {
	if s.head <= s.tail {
		return s.tail - s.head
	} else {
		return s.size - s.head + s.tail
	}
}

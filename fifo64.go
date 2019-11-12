package fifo64

type FifoItem uint64

type Fifo struct {
	head int
	tail int
	data []FifoItem
	size int
}

func New(size int) *Fifo {
	f := new(Fifo)
	f.data = make([]FifoItem, size+1, size+1)
	f.size = size
	f.head = 0
	f.tail = 0
	return f
}

func (f *Fifo) Len() int {
	if f.head <= f.tail {
		return f.tail - f.head
	} else {
		return f.size - f.head + f.tail
	}
}

// I assume that this API is "reasonably" tread safe. Will not cause
// problems if there is a race
// s.head is modified by remove() and is an atomic operation
// I do not care about valifity of s.tai
func (f *Fifo) pick() (key uint64, ok bool) {
	if f.head != f.tail {
		key = f.data[f.head]
		return key, true
	} else {
		return key, false
	}
}


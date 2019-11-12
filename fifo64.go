package fifo64

type Fifo struct {
	head int
	tail int
	data []FifoItem
	size int
}

func New(size int) *Fifo {
	f := new(Fifo)
	f.data = make([]uint64, size+1, size+1)
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
func (f *Fifo) Pick() (key uint64, ok bool) {
	if f.head != f.tail {
		key = f.data[f.head]
		return key, true
	} else {
		return key, false
	}
}


func (f *Fifo) add(key uint64) (ok bool) {
	newTail := s.inc(s.tail)
	if s.head != newTail {
		s.data[s.tail] = key
		s.tail = newTail
		return true
	} else {
		return false
	}
}

func (f *Fifo) remove() (key uint64, ok bool) {
	newHead := s.inc(s.head)
	if s.head != s.tail {
		key = s.data[s.head]
		s.head = newHead
		return key, true
	} else {
		return key, false
	}
}

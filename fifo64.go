package fifo64

// Fifo of uint64 items
type Fifo struct {
	head int
	tail int
	data []uint64
	size int
}

// New returns a new FIFO
func New(size int) *Fifo {
	f := new(Fifo)
	f.data = make([]uint64, size+1, size+1)
	f.size = size
	f.head = 0
	f.tail = 0
	return f
}

// Len returns occupancy
func (f *Fifo) Len() int {
	if f.head <= f.tail {
		return f.tail - f.head
	}
	return f.size - f.head + f.tail
}

// Pick an item from the head
// I assume that this API is "reasonably" tread safe. Will not cause
// problems if there is a race
// s.head is modified by remove() and is an atomic operation
// I do not care about valifity of s.tai
func (f *Fifo) Pick() (key uint64, ok bool) {
	if f.head != f.tail {
		key = f.data[f.head]
		return key, true
	}
	return key, false
}

// Add an item to the FIFO
func (f *Fifo) Add(key uint64) (ok bool) {
	newTail := f.inc(f.tail)
	if f.head != newTail {
		f.data[f.tail] = key
		f.tail = newTail
		return true
	}
	return false
}

// Remove an item from the FIFO
func (f *Fifo) Remove() (key uint64, ok bool) {
	newHead := f.inc(f.head)
	if f.head != f.tail {
		key = f.data[f.head]
		f.head = newHead
		return key, true
	}
	return key, false
}

func (f *Fifo) inc(v int) int {
	if v < f.size {
		v++
	} else {
		v = 0
	}
	return v
}

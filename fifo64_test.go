package fifo64

import (
	"testing"
)

func BenchmarkFifo(b *testing.B) {
	fifoSize := 10 * 1000 * 1000
	fifo := New(fifoSize)
	b.ReportAllocs()
	b.N = fifoSize
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ok := fifo.Add(uint64(i))
		if !ok {
			b.Fatalf("Failed to add an object to the FIFO %d", i)
		}
	}
}

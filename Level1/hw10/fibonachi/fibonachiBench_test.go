package fibonachi

import "testing"

func BenchmarkFibonachi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonachi(30)
	}
}

func BenchmarkFibonachiCached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonachiCached(30)
	}
}

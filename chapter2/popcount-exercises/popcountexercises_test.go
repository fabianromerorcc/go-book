package popcountexercises

import "testing"

var x uint64 = 4294967296

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(x)
	}
}

func BenchmarkPopCountClearRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearRight(x)
	}
}

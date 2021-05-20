package fmath

import (
	"testing"
)

func TestSqrtf(t *testing.T) {
	f := Sqrtf(2)
	if f != 1.41421356237 {
		t.Fatal("Sqrtf(2):", f)
	}
}

func BenchmarkSqrtfAssembly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sqrtf(2)
	}
}

func BenchmarkSqrtfGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sqrt(2)
	}
}
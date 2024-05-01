package fmath

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	f := Sqrt(2)
	if f != 1.41421356237 {
		t.Fatal("Sqrt(2):", f)
	}
}

func BenchmarkSqrtAssembly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sqrt(2)
	}
}

func BenchmarkSqrtGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sqrt(2)
	}
}

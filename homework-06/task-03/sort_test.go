package main

import (
	"testing"
	"sort"
)

func BenchmarkInts(b *testing.B) {
	s := []int{5, 2, 6, 3, 1, 4}

	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}

func BenchmarkFloat64s(b *testing.B) {
	s := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}

	for i := 0; i < b.N; i++ {
		sort.Float64s(s)
	}
}
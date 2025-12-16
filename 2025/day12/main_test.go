package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	gifts, regions := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partOne(gifts, regions)
	}
}
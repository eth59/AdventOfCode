package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partOne(input, 101, 103)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partTwo(input, 101, 103)
	}
}
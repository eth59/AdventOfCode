package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	inputGrid := parseGrid(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partOne(inputGrid)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	inputGrid := parseGrid(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partTwo(inputGrid)
	}
}
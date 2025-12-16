package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	devices := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partOne(devices)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	devices := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partTwo(devices)
	}
}
package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	machineInput := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partOne(machineInput)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	machineInput := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partTwo(machineInput)
	}
}
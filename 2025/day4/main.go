package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input_test.txt
var inputTest string

//go:embed input.txt
var input string

func main() {
	inputTestGrid := parseGrid(inputTest)
	inputGrid := parseGrid(input)

	fmt.Println("PART 1:")

	resultTest := partOne(inputTestGrid)
	fmt.Println("Result test:", resultTest)

	result := partOne(inputGrid)
	fmt.Println("Result:", result)

	fmt.Println("PART 2:")

	resultTest2 := partTwo(inputTestGrid)
	fmt.Println("Result test:", resultTest2)

	result = partTwo(inputGrid)
	fmt.Println("Result:", result)
}

func parseGrid(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for r, line := range lines {
		grid[r] = []rune(line)
	}
	return grid
}
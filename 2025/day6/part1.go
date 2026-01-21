package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	grid := parseInput(input)
	height := len(grid)
	width := len(grid[0])

	for c := 0; c < width; c++ {
		op := grid[height-1][c]
		columnRes, _ := strconv.Atoi(grid[0][c])
		for r := 1; r < height-1; r++ {
			current, _ := strconv.Atoi(grid[r][c])
			if op == "+" {
				columnRes += current
			} else {
				columnRes *= current
			}
		}
		res += columnRes
	}
	return
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, 0, len(lines))
	for _, line := range lines {
		grid = append(grid, strings.Fields(line))
	}
	return grid
}
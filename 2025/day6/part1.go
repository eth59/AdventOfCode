package main

import (
	"strconv"
	"strings"
)

func partOne(input string) int {
	grid := parseInput(input)
	res := 0
	height := len(grid)
	width := len(grid[0])

	for c := 0; c < width; c++ {
		op := grid[height-1][c]
		columnRes, err := strconv.Atoi(grid[0][c])
		if err != nil {
			panic(err)
		}
		for r := 1; r < height-1; r++ {
			current, err := strconv.Atoi(grid[r][c])
			if err != nil {
				panic(err)
			}
			if op == "+" {
				columnRes += current
			} else {
				columnRes *= current
			}
		}
		res += columnRes
	}
	return res
}

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]string, 0, len(lines))
	for _, line := range lines {
		grid = append(grid, strings.Fields(line))
	}
	return grid
}
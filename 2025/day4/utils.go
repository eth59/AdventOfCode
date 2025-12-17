package main

import "strings"

func parseGrid(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for r, line := range lines {
		grid[r] = []rune(line)
	}
	return grid
}
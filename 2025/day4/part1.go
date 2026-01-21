package main

import (
	"strings"
)

var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, 		   {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func partOne(input string) (res int) {
	grid := parseGrid(input)

	height := len(grid)
	width := len(grid[0])

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '@' {
				nbNeighbors := 0
				for _, dir := range directions {
					nr, nc := r+dir[0], c+dir[1]
					if nr >= 0 && nr < height && nc >= 0 && nc < width && grid[nr][nc] == '@' {
						nbNeighbors++
					}
				}
				if nbNeighbors < 4 {
					res++
				}
			}
		}
	}
	return
}

func parseGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	for r, line := range lines {
		grid[r] = []rune(line)
	}
	return grid
}
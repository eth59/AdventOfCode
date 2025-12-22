package main

import (
	"strings"
)

type Point struct {
	R, C int // Row, Col
}

var dirs = []Point{
	{-1, 0}, // 0 : haut
	{0, 1}, // 1 : droite
	{1, 0}, // 2 : bas
	{0, -1}, // 3 : gauche
}

func (p Point) Add(p2 Point) Point {
	return Point{p.R + p2.R, p.C + p2.C}
}

func partOne(input string) (res int) {
	grid := strings.Split(input, "\n")
	height := len(grid)
	width := len(grid[0])

	// on cherche le garde
	guardPos := findGuard(grid, height, width)
	currentDirIdx := 0

	// on parcourt la grille en suivant le garde et en comptant le nb de cases visitées
	visited := make(map[Point]bool)
	for {
		newPos := guardPos.Add(dirs[currentDirIdx])
		if newPos.R < 0 || newPos.R >= height || newPos.C < 0 || newPos.C >= height {
			return
		}
		if grid[newPos.R][newPos.C] == '#' {
			currentDirIdx = (currentDirIdx + 1) % 4
		} else {
			guardPos = newPos
			if !visited[guardPos] {
				visited[guardPos] = true
				res ++
			}
		}
	}
}

func findGuard(grid []string, height, width int) Point {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '^' {
				return Point{r, c}
			}
		}
	}
	return Point{}
}
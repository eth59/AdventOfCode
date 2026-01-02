package main

import (
	"strings"
)

func partTwo(input string) (res int) {
	// parsing
	grid := strings.Split(input, "\n")
	height := len(grid)
	width := len(grid[0])

	// variables
	visited := make(map[Point]bool)

	// parcours
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			p := Point{r, c}

			if visited[p] {
				continue
			}

			area, sides := exploreRegionTwo(grid, p, visited, height, width)
			res += area * sides
		}
	}

	return
}

func exploreRegionTwo(grid []string, p Point, visited map[Point]bool, h, w int) (area, sides int) {
	letter := grid[p.R][p.C]
	queue := []Point{p}
	visited[p] = true

	for len(queue) > 0 {
		// on dépile
		curr := queue[0]
		queue = queue[1:]

		area ++ // on a trouvé une nouvelle case dans la région
		sides += countCorners(grid, curr, letter, h, w) // nb côtés = nb coins

		for _, d := range dirs {
			neighbour := Point{curr.R + d.R, curr.C + d.C}

			if neighbour.R >= 0 && neighbour.R < h && neighbour.C >= 0 && neighbour.C < w && grid[neighbour.R][neighbour.C] == letter {
				// nouvelle case de la région
				if !visited[neighbour] {
					visited[neighbour] = true
					queue = append(queue, neighbour)
				}
			}
		}
	}

	return
}

func countCorners(grid []string, p Point, letter byte, h, w int) (nbCorners int) {
	sameLetter := func(r, c int) bool {
		return r >= 0 && r < h && c >= 0 && c < w && grid[r][c] == letter
	}

	// check si même lettre dans les quatres directions
	up := sameLetter(p.R-1, p.C)
	right := sameLetter(p.R, p.C+1)
	down := sameLetter(p.R+1, p.C)
	left := sameLetter(p.R, p.C-1)

	// check si même lettre dans les diags
	ul := sameLetter(p.R-1, p.C-1)
	ur := sameLetter(p.R-1, p.C+1)
	dr := sameLetter(p.R+1, p.C+1)
	dl := sameLetter(p.R+1, p.C-1)

	// coin haut gauche
	if (!up && !left) || (up && left && !ul) { nbCorners++ }

	// coin haut droite
	if (!up && !right) || (up && right && !ur) { nbCorners++ }

	// coin bas droite
	if (!down && !right) || (down && right && !dr) { nbCorners++ }

	// coin bas gauche
	if (!down && !left) || (down && left && !dl) { nbCorners++ }
	
	return
}
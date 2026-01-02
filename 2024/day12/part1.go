package main

import "strings"

type Point struct {
	R, C int
}

type Region struct {
	Area, Perimeter int
}

var dirs = []Point{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func partOne(input string) (res int) {
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

			area, perimeter := exploreRegion(grid, p, visited, height, width)
			res += area * perimeter
		}
	}

	return
}

func exploreRegion(grid []string, p Point, visited map[Point]bool, h, w int) (area, perimeter int) {
	letter := grid[p.R][p.C]
	queue := []Point{p}
	visited[p] = true

	for len(queue) > 0 {
		// on dépile
		curr := queue[0]
		queue = queue[1:]

		area ++ // on a trouvé une nouvelle case dans la région

		for _, d := range dirs {
			neighbour := Point{curr.R + d.R, curr.C + d.C}

			if neighbour.R >= 0 && neighbour.R < h && neighbour.C >= 0 && neighbour.C < w && grid[neighbour.R][neighbour.C] == letter {
				// nouvelle case de la région
				if !visited[neighbour] {
					visited[neighbour] = true
					queue = append(queue, neighbour)
				}
			} else {
				// c'est un bord
				perimeter ++
			}
		}
	}

	return
}
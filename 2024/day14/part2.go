package main

import (
	"fmt"
	"strings"
)

type Robot struct {
	X, Y, Vx, Vy int
}

func partTwo(input string, w, h int) int {
	var robots []Robot

	// parsing
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var r Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.X, &r.Y, &r.Vx, &r.Vy)
		robots = append(robots, r)
	}

	var t = 1
	for {
		// on va faire une grille pour la seconde actuelle
		grid := make([][]bool, h)
		for i := 0; i < h; i++ {
			grid[i] = make([]bool, w)
		}

		// on déplace les robots et met à true dans la grille
		for i := range robots {
			robots[i].X = ((robots[i].X+robots[i].Vx) % w + w) % w
			robots[i].Y = ((robots[i].Y+robots[i].Vy) % h + h) % h
			grid[robots[i].Y][robots[i].X] = true
		}

		// heuristique : ligne de 8 robots consécutifs avec affichage facultatif
		if hasLine(grid, w, h, 8) {
			// printGrid(grid, w, h)
			return t
		}
		t++
	}
}

// vérifie s'il y a une ligne de x robots
func hasLine(grid [][]bool, w, h, x int) bool {
	for r := 0; r < h; r++ {
		consecutive := 0
		for c := 0; c < w; c++ {
			if grid[r][c] {
				consecutive++
				if consecutive >= x {
					return true
				}
			} else {
				consecutive = 0
			}
		}
	}
	return false
}

// affiche la grille pour voir le sapin
func printGrid(grid [][]bool, w, h int) {
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] {
				fmt.Print("#")
			} else {
				fmt.Print(".")		
			}
		}
		fmt.Println()
	}
}
package main

import "strings"

func partTwo(input string) (res int) {
	grid := strings.Split(input, "\n")
	gridSize := len(grid) // la grille est carrée

	// on parcourt la grille et à chaque 0, on lance la fonction de solve
	for r, line := range grid {
		for c, digit := range line {
			if digit == '0' {
				res += solvePartTwo(grid, gridSize, Point{r, c}, 0)
			}
		}
	}
	return
}

// on renvoie la liste des 9 accessible
func solvePartTwo(grid []string, gridSize int, p Point, currentHeight int) (res int) {
	// condition d'arrêt, on a trouvé un chemin
	if currentHeight == 9 {
		return 1
	}
	// on regarde en haut si possible
	if p.R > 0 && int(grid[p.R-1][p.C] - '0') == currentHeight + 1 {
		res += solvePartTwo(grid, gridSize, Point{p.R-1, p.C}, currentHeight+1)
	}
	// on regarde à droite si possible
	if p.C < gridSize-1 && int(grid[p.R][p.C+1] - '0') == currentHeight + 1 {
		res += solvePartTwo(grid, gridSize, Point{p.R, p.C+1}, currentHeight+1)
	}
	// on regarde en bas si possible
	if p.R < gridSize-1 && int(grid[p.R+1][p.C] - '0') == currentHeight + 1 {
		res += solvePartTwo(grid, gridSize, Point{p.R+1, p.C}, currentHeight+1)
	}
	// on regarde à gauche si possible
	if p.C > 0 && int(grid[p.R][p.C-1] - '0') == currentHeight + 1 {
		res += solvePartTwo(grid, gridSize, Point{p.R, p.C-1}, currentHeight+1)
	}
	return
}
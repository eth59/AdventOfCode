package main

import "strings"

type PointAndDir struct {
	P Point
	dirIdx int
}

func partTwo(input string) (res int) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])

	// on convertit la grille en [][]rune pour pouvoir la modifier & on trouve le garde
	grid := make([][]rune, height)
	var guardPos Point
	for r, line := range lines {
		grid[r] = []rune(line)
		if c := strings.IndexRune(line, '^'); c != -1 {
			guardPos = Point{r, c}
		}
	}

	// on parcourt les points du chemin original, on essaye d'y mettre un objet et on voit si ça boucle
	for _, p := range findPath(grid, guardPos, height, width) {
		// on modifie la grille
		grid[p.R][p.C] = '#'

		visited := make(map[PointAndDir]bool) // on garde aussi la direction, pour trouver une boucle
		// si on revient à la même position avec la même direction, on boucle
		currentGuardPos := guardPos // on veut pas modifier la position originale du garde
		// on parcourt le chemin du garde à la recherche de boucle ou non
		currentDirIdx := 0
		for {
			newPos := currentGuardPos.Add(dirs[currentDirIdx])
			if newPos.R < 0 || newPos.R >= height || newPos.C < 0 || newPos.C >= width {
				break // on est sorti de la map sans boucler
			}
			if grid[newPos.R][newPos.C] == '#' {
				currentDirIdx = (currentDirIdx + 1) % 4
			} else {
				currentGuardPos = newPos
				currentPosAndDir := PointAndDir{currentGuardPos, currentDirIdx}
				if !visited[currentPosAndDir] {
					visited[currentPosAndDir] = true
				} else {
					// on a bouclé
					res++
					break
				}
			}
		}

		grid[p.R][p.C] = '.' // on restaure la map
	}

	return
}

// find the original path
func findPath(grid [][]rune, guardPos Point, height, width int) (path []Point) {
	visited := make(map[Point]bool)
	path = make([]Point, 0, width*height)
	currentDirIdx := 0
	for {
		newPos := guardPos.Add(dirs[currentDirIdx])
		if newPos.R < 0 || newPos.R >= height || newPos.C < 0 || newPos.C >= width {
			return
		}
		if grid[newPos.R][newPos.C] == '#' {
			currentDirIdx = (currentDirIdx + 1) % 4
		} else {
			guardPos = newPos
			if !visited[guardPos] {
				visited[guardPos] = true
				path = append(path, guardPos)
			}
		}
	}
}
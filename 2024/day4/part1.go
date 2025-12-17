package main

import "strings"

func partOne(input string) (res int) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	var up, right, down, left bool // pour stocker dans quel sens c possible (déborde pas du word search)
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			// on regarde dans quel sens ça ne déborde pas
			up = r > 2
			right = c < width-3
			down = r < height-3
			left = c > 2

			// on lance la recherche dans les 8 directions si ça ne déborde pas
			if up && searchXMAS(lines, r, c, -1, 0){
				// vers le haut
				res ++
			}
			if up && right && searchXMAS(lines, r, c, -1, 1) {
				// vers le haut à droite
				res++
			}
			if right && searchXMAS(lines, r, c, 0, 1) {
				// vers la droite
				res++
			}
			if down && right && searchXMAS(lines, r, c, 1, 1) {
				// vers en bas à droite
				res++
			}
			if down && searchXMAS(lines, r, c, 1, 0) {
				// vers le bas
				res++
			}
			if down && left && searchXMAS(lines, r, c, 1, -1) {
				// vers en bas à gauche
				res++
			}
			if left && searchXMAS(lines, r, c, 0, -1) {
				// vers la gauche
				res++
			}
			if left && up && searchXMAS(lines, r, c, -1, -1) {
				// vers en haut à gauche
				res++
			}
		}
	}
	return
}

func searchXMAS(lines []string, x, y, dx, dy int) bool {
	if lines[x][y] != 'X' {
		return false
	} else if lines[x+dx][y+dy] != 'M' {
		return false
	} else if lines[x+2*dx][y+2*dy] != 'A' {
		return false
	} else if lines[x+3*dx][y+3*dy] != 'S' {
		return false
	}
	return true
}
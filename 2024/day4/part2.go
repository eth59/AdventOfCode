package main

import (
	"strings"
)

func partTwo(input string) (res int) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	for r := 1; r < height-1; r++ {
		for c := 1; c < width-1; c++ {
			// ici [r][c] = 'A' comparé à la part 1
			if isXMAS(lines, r, c) {
				res++
			}
		}
	}
	return
}

func isXMAS(lines []string, r, c int) bool {
	// on veut un A au centre
	if lines[r][c] != 'A' {
		return false
	}

	// on récupère les 4 caractères dans les coins
	ul := lines[r-1][c-1]
	ur := lines[r-1][c+1]
	dr := lines[r+1][c+1]
	dl := lines[r+1][c-1]

	// on vérifie les deux diags
	diag1 := (ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M') // \
	diag2 := (ur == 'M' && dl == 'S') || (ur == 'S' && dl == 'M') // /
	return diag1 && diag2
}
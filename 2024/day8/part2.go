package main

import (
	"strings"
)

func partTwo(input string) (res int) {
	// parse
	antennas := make(map[rune][]Point)
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	for r, line := range lines {
		for c, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], Point{r, c})
			}
		}
	}

	// calcul du res
	validAntinodes := make(map[Point]bool)
	for _, antennasPoints := range antennas {
		length := len(antennasPoints)
		for i := 0; i < length; i++ {
			for j := i+1; j < length; j++ {
				// on récupére les deux antennes
				a1 := antennasPoints[i]
				a2 := antennasPoints[j]

				// on calcule les deux antinodes
				dist := sub(a2, a1)
				
				// vérif & ajout
				curr := a1
				for isInBounds(curr, height, width) {
					validAntinodes[curr] = true
					curr = sub(curr, dist)
				}
				curr = a2
				for isInBounds(curr, height, width) {
					validAntinodes[curr] = true
					curr = add(curr, dist)
				}
			}
		}
	}

	return len(validAntinodes)
}
package main

import "strings"

type Point struct {
	R, C int // Row, Col
}

func add(p1, p2 Point) Point {
	return Point{p1.R+p2.R, p1.C+p2.C} // p1 + p2
}

func sub(p1, p2 Point) Point {
	return Point{p1.R-p2.R, p1.C-p2.C} // p1 - p2
}

func partOne(input string) int {
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
				anti1 := sub(a1, dist)
				anti2 := add(a2, dist)
				
				// vérif & ajout
				if isInBounds(anti1, height, width) {
					validAntinodes[anti1] = true
				}
				if isInBounds(anti2, height, width) {
					validAntinodes[anti2] = true
				}
			}
		}
	}

	return len(validAntinodes)
}

func isInBounds(antinode Point, height, width int) bool {
	return antinode.R >= 0 && antinode.R < height && antinode.C >= 0 && antinode.C < width
}
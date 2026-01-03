package main

import (
	"fmt"
	"strings"
)

func partOne(input string, w, h int) int {
	var quadrants [4]int // 0 : haut gauche, 1 : haut droite, 2 : bas droite, 3 : bas gauche
	midX := w/2
	midY := h/2

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// parsing
		var x, y, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)

		// move robot 100 times (on force un res positif aux modulos)
		x = ((x+100*vx) % w + w) % w
		y = ((y+100*vy) % h + h) % h

		// calcul res selon quart
		if x < midX && y < midY {
			quadrants[0]++
		} else if x > midX && y < midY {
			quadrants[1]++
		} else if x > midX && y > midY {
			quadrants[2]++
		} else if x < midX && y > midY {
			quadrants[3]++
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}
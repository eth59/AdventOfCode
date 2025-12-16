package main

import (
	"math"
	"strconv"
	"strings"
)

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	nbRedTiles := len(lines)

	// parse
	redTiles := make([]coords, 0, nbRedTiles)
	for _, line := range lines {
		c := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(c[0], 64)
		y, _ := strconv.ParseFloat(c[1], 64)
		redTiles = append(redTiles, coords{x, y})
	}

	// find max area
	var maxArea float64
	var newArea float64
	for i := 0; i < nbRedTiles; i++ {
		for j := i+1; j < nbRedTiles; j++ {
			newArea = (math.Abs(float64(redTiles[i].x) - float64(redTiles[j].x)) + 1) * (math.Abs(float64(redTiles[i].y) - float64(redTiles[j].y)) + 1)
			if newArea > maxArea {
				maxArea = newArea
			}
		}
	}

	return int(maxArea)
}
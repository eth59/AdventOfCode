package main

import (
	"fmt"
	"sort"
	"strings"
)


func partTwo(input string, width, height int) string {
	lines := strings.Split(input, "\n")
	corruptedPoints := make([]Point, 0, len(lines))
	var x, y int
	for _, line := range lines {
		fmt.Sscanf(line, "%d,%d", &x, &y)
		corruptedPoints = append(corruptedPoints, Point{x, y})
	}

	idx := sort.Search(len(corruptedPoints), func(i int) bool {
		corruptedPointsToTest := make([][]bool, height)
		for k := 0; k < height; k++ {
			corruptedPointsToTest[k] = make([]bool, width)
		}

		for k := 0; k <= i; k++ {
			corruptedPointsToTest[corruptedPoints[k].y][corruptedPoints[k].x] = true
		}

		return findPath(corruptedPointsToTest, width, height) == -1
	})

	return fmt.Sprintf("%d,%d", corruptedPoints[idx].x, corruptedPoints[idx].y)
}
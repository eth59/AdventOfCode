package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

type State struct {
	point Point
	dist int
}

// 0 : droite, 1 : bas, 2 : gauche, 3 : haut
var dirs = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func partOne(input string, nbBytesCorrupted, width, height int) int {
	corruptedPoints := parse(input, nbBytesCorrupted, width, height)
	return findPath(corruptedPoints, width, height)
}

func findPath(corruptedPoints [][]bool, width, height int) int {
	startingPoint := Point{0, 0}
	endingPoint := Point{width-1, height-1}

	visited := make(map[Point]bool, width*height) // pour pas tourner en rond
	visited[startingPoint] = true

	queue := make([]State, 0)
	queue = append(queue, State{startingPoint, 0})

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.point == endingPoint {
			return curr.dist
		}

		for _, dir := range dirs {
			nx, ny := curr.point.x + dir.x, curr.point.y + dir.y
			next := Point{nx, ny}

			if nx >= 0 && nx < width && ny >= 0 && ny < height && !visited[next] && !corruptedPoints[ny][nx] {
				visited[next] = true
				queue = append(queue, State{next, curr.dist+1})
			}
		}
	}

	return -1 // normalement c'est impossible
}

func parse(input string, nbBytesCorrupted, width, height int) (corruptedPoints [][]bool) {
	lines := strings.Split(input, "\n")
	length := len(lines)
	corruptedPoints = make([][]bool, height)
	for i := 0; i < height; i++ {
		corruptedPoints[i] = make([]bool, width)
	}

	var x, y int

	for i := 0; i < nbBytesCorrupted && i < length; i++ {
		fmt.Sscanf(lines[i], "%d,%d", &x, &y)
		corruptedPoints[y][x] = true
	}

	return
}
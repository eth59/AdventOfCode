package main

import (
	"sort"
	"strconv"
	"strings"
)

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	nbBoxes := len(lines)
	junctionBoxes := make([]coordinates, 0, nbBoxes)

	// parse boxes
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			panic(err)
		}
		junctionBoxes = append(junctionBoxes, coordinates{x, y, z})
	}

	// init circuits
	circuits = make([]int, nbBoxes)
	size = make([]int, nbBoxes)
	for i := 0; i < nbBoxes; i++ {
		circuits[i] = i
		size[i] = 1
	}

	// compute distances
	distances := make([]Pair, 0, nbBoxes*(nbBoxes-1)/2)
	for i := 0; i < nbBoxes; i++ {
		for j := i+1; j < nbBoxes; j++ {
			distances = append(distances, Pair{computeDistance(junctionBoxes[i], junctionBoxes[j]), i, j})
		}
	}

	// tri par distance
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	// on fait les connections
	remainingCircuits := nbBoxes

	for _, p := range distances {
		if union(p.i, p.j) {
			remainingCircuits--
			if remainingCircuits == 1 {
				return junctionBoxes[p.i].x * junctionBoxes[p.j].x
			}
		}
	}

	return -1
}
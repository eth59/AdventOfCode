package main

import (
	"sort"
	"strconv"
	"strings"
)

type coordinates struct {
	x, y, z int
}

type Pair struct {
	dist int
	i int // index dans slice junctionBoxes
	j int // same
}

var circuits []int
var size []int

func partOne(input string, nbPairs int) (res int) {
	lines := strings.Split(input, "\n")
	nbBoxes := len(lines)
	junctionBoxes := make([]coordinates, 0, nbBoxes)

	// parse boxes
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
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
	for i := 0; i < nbPairs && i < len(distances); i++ {
        p := distances[i]
        union(p.i, p.j)
    }
	// compute taille des circuits
	circuitsSizes := make([]int, 0)
	for i := 0; i < nbBoxes; i++ {
		if circuits[i] == i {
			circuitsSizes = append(circuitsSizes, size[i])
		}
	}

	// compute res
	sort.Slice(circuitsSizes, func(i, j int) bool {
		return circuitsSizes[i] > circuitsSizes[j]
	})

	res = circuitsSizes[0] * circuitsSizes[1] * circuitsSizes[2]

	return
}

func computeDistance(a, b coordinates) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func find(i int) int {
	if circuits[i] == i {
		return i
	}
	circuits[i] = find(circuits[i])
	return circuits[i]
}

func union(i, j int) bool {
	rootI, rootJ := find(i), find(j)

	if rootI != rootJ {
		if size[rootI] < size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		circuits[rootJ] = rootI
		size[rootI] += size[rootJ]

		return true
	}

	return false
}
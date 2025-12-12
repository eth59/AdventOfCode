package main

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
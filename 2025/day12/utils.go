package main

import (
	"strconv"
	"strings"
)

type Gift struct {
	shape [9]bool
	Area int
}

func (g Gift) set(x, y int, value bool) Gift {
	g.shape[x*3+y] = value
	return g
}

func (g Gift) get(x, y int) bool {
	return g.shape[x*3+y]
}

// String pour le debug
func (g Gift) String() (res string) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if g.get(x, y) {
				res += "#"
			} else {
				res += "."
			}
		}
		if x < 2 {
			res += "\n"
		}
	}
	return
}

// rotate dans le sens horaire
func (g Gift) rotate() (res Gift) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			res = res.set(x, y, g.get(2-y, x))
		}
	}
	return
}

// flip horizontal
func (g Gift) flip() (res Gift) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			res = res.set(x, y, g.get(x, 2-y))
		}
	}
	return
}

type Region struct {
	SizeX, SizeY int
	NeededShapes []int
	NeededGifts int
}

func parseInput(input string) (gifts []Gift, regions []Region) {
	parts := strings.Split(input, "\n\n")

	// parse gifts
	for _, part := range parts[:len(parts)-1] {
		gift := Gift{}
		area := 0
		lines := strings.Split(part, "\n")
		for i, line := range lines[1:] {
			for j, c := range line {
				if c == '#' {
					area++
					gift = gift.set(i, j, true)
				} else {
					gift = gift.set(i, j, false)
				}
			}
		}
		gift.Area = area
		gifts = append(gifts, gift)
	}

	// parse regions
	for _, line := range strings.Split(parts[len(parts)-1], "\n") {
		parts := strings.Split(line, ": ")
		size := strings.Split(parts[0], "x")
		sizeX, _ := strconv.Atoi(size[0])
		sizeY, _ := strconv.Atoi(size[1])
		neededShapesStr := strings.Split(parts[1], " ")
		neededShapes := make([]int, 0, len(neededShapesStr))
		var neededGifts int
		for _, nbStr := range neededShapesStr {
			nb, _ := strconv.Atoi(nbStr)
			neededShapes = append(neededShapes, nb)
			neededGifts += nb
		}

		regions = append(regions, Region{sizeX, sizeY, neededShapes, neededGifts})
	}

	return
}
package main

import (
	"fmt"
	"strings"
)

var MAX_RED_CUBES = 12
var MAX_GREEN_CUBES = 13
var MAX_BLUE_CUBES = 14

func partOne(input string) (res int) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		cubesSets := strings.Split(strings.ReplaceAll(strings.Split(line, ": ")[1], "; ", ", "), ", ") 

		var nbCubes int
		var cubesColor string
		isPossible := true

		for _, cubesSet := range cubesSets {
			fmt.Sscanf(cubesSet, "%d %s", &nbCubes, &cubesColor)

			switch cubesColor {
			case "red":
				if nbCubes > MAX_RED_CUBES {
					isPossible = false
				}
			case "green":
				if nbCubes > MAX_GREEN_CUBES {
					isPossible = false
				}
			case "blue":
				if nbCubes > MAX_BLUE_CUBES {
					isPossible = false
				}
			}
		}

		if isPossible {
			res += i + 1
		}
	}
	return
}
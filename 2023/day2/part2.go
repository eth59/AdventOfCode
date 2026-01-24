package main

import (
	"fmt"
	"strings"
)

func partTwo(input string) (res int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		cubesSets := strings.Split(strings.ReplaceAll(strings.Split(line, ": ")[1], "; ", ", "), ", ") 

		var minNbRed, minNbGreen, minNbBlue, nbCubes int
		var cubesColor string

		for _, cubesSet := range cubesSets {
			fmt.Sscanf(cubesSet, "%d %s", &nbCubes, &cubesColor)

			switch cubesColor {
			case "red":
				if nbCubes > minNbRed {
					minNbRed = nbCubes
				}
			case "green":
				if nbCubes > minNbGreen {
					minNbGreen = nbCubes
				}			
			case "blue":
				if nbCubes > minNbBlue {
					minNbBlue = nbCubes
				}
			}
		}

		res += minNbRed * minNbGreen * minNbBlue
	}
	return
}
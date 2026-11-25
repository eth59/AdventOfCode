package main

import (
	"strconv"
	"strings"
)

func partTwo(input string) int {
	return solveOneRace(parsePartTwo(input))
}

func parsePartTwo(input string) (float64, float64) {
	lines := strings.Split(input, "\n")

	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	timeStr := ""
	distanceStr := ""

	for i, t := range times {
		timeStr += t
		distanceStr += distances[i]
	}
	
	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	return float64(time), float64(distance)
}
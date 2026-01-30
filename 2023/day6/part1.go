package main

import (
	"math"
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	res = 1
	times, distances, length := parsePartOne(input)
	for i := 0; i < length; i++ {
		tmp := solveOneRace(times[i], distances[i])
		res *= tmp
	}
	return
}

func solveOneRace(time float64, distance float64) int {
	discriminant := time*time - 4*distance

	epsilon := 1e-9 // sans espilon, on peut avoir des solutions entières et "compter" une valeur de trop

	minSol := (time - math.Sqrt(discriminant))/2 + epsilon
	maxSol := (time + math.Sqrt(discriminant))/2 - epsilon

	return int(math.Floor(maxSol) - math.Floor(minSol))
}

func parsePartOne(input string) (times []float64, distances []float64, length int) {
	lines := strings.Split(input, "\n")

	// times
	timesStr := strings.Fields(lines[0])
	length = len(timesStr)-1
	times = make([]float64, length)
	for i, timeStr := range timesStr[1:] {
		time, _ := strconv.Atoi(timeStr)
		times[i] = float64(time)
	}

	// distances
	distancesStr := strings.Fields(lines[1])
	distances = make([]float64, length)
	for i, distanceStr := range distancesStr[1:] {
		distance, _ := strconv.Atoi(distanceStr)
		distances[i] = float64(distance)
	}

	return
}
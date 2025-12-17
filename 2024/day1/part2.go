package main

import (
	"strconv"
	"strings"
)

func partTwo(input string) (res int) {
	// parsing
	lines := strings.Split(input, "\n")
	// map plutôt que les slices de la part 1
	occurencesLeft := make(map[int]int)
	occurencesRight := make(map[int]int)
	for _, line := range lines {
		nbs := strings.Split(line, "   ")
		nbLeft, _ := strconv.Atoi(nbs[0])
		nbRight, _ := strconv.Atoi(nbs[1])
		occurencesLeft[nbLeft]++
		occurencesRight[nbRight]++
	}

	// calcul res
	for nb, occ := range occurencesLeft {
		res += occ * nb * occurencesRight[nb]
	}
	return
}
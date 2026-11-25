package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		goal, _ := strconv.Atoi(parts[0])
		nbsStr := strings.Fields(parts[1])
		nbsLength := len(nbsStr)
		nbs := make([]int, 0, nbsLength)
		for _, nbStr := range nbsStr {
			nb, _ := strconv.Atoi(nbStr)
			nbs = append(nbs, nb)
		}
		if solvePartOne(goal, nbs, nbs[0], 1, nbsLength) {
			res += goal
		}
	}
	return
}

func solvePartOne(goal int, nbs []int, currentNb, idx, length int) bool {
	if currentNb == goal && idx == length {
		return true // on a trouvé une solution
	}
	if currentNb > goal || idx >= length {
		return false // les opérations permettent seulement d'augmenter ou index dépassé
	}
	return solvePartOne(goal, nbs, currentNb+nbs[idx], idx+1, length) || solvePartOne(goal, nbs, currentNb*nbs[idx], idx+1, length)
}
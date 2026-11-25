package main

import (
	"strings"
	"strconv"
)

func partTwo(input string) (res int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		goal, _ := strconv.Atoi(parts[0])
		nbsStr := strings.Fields(parts[1]) // on garde sous forme de string pour faciliter la conversion
		nbsLength := len(nbsStr)
		currentNb, _ := strconv.Atoi(nbsStr[0])
		if solvePartTwo(goal, nbsStr, currentNb, 1, nbsLength) {
			res += goal
		}
	}
	return
}

func solvePartTwo(goal int, nbs []string, currentNb int, idx, length int) bool {
	if currentNb == goal && idx == length {
		return true // on a trouvé une solution
	}
	if currentNb > goal || idx >= length {
		return false // les opérations permettent seulement d'augmenter ou index dépassé
	}
	nbAtIdx, _ := strconv.Atoi(nbs[idx])
	concatNb, _ := strconv.Atoi(strconv.Itoa(currentNb) + nbs[idx]) 
	return (solvePartTwo(goal, nbs, currentNb+nbAtIdx, idx+1, length) ||
			solvePartTwo(goal, nbs, currentNb*nbAtIdx, idx+1, length) ||
			solvePartTwo(goal, nbs, concatNb, idx+1, length))
}
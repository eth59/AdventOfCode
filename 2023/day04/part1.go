package main

import (
	"slices"
	"strings"
)

func partOne(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		nbWinningNbsInList := getNbWinningNumbers(line)
		if nbWinningNbsInList > 0 {
			res += 1 << (nbWinningNbsInList - 1)
		}
	}
	return
}

func getNbWinningNumbers(line string) (nbWinningNbsInList int) {
	parts := strings.Split(strings.Split(line, ": ")[1], " | ")
	winningNbs := strings.Fields(parts[0])
	haveNbs := strings.Fields(parts[1])
	for _, nb := range haveNbs {
		if slices.Contains(winningNbs, nb) {
			nbWinningNbsInList++
		}
	}
	return
}
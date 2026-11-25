package main

import "strings"

func partTwo(input string) (res int) {
	lines := strings.Split(input, "\n")
	nbCards := len(lines)
	nbCopies := make([]int, nbCards)
	for i, line := range lines {
		nbCopies[i]++
		nbWinningNbsInList := getNbWinningNumbers(line)
		for j := 1; j <= nbWinningNbsInList; j++ {
			nbCopies[i+j] += nbCopies[i]
		}
	}
	for _, n := range nbCopies {
		res += n
	}
	return
}
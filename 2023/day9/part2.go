package main

import (
	"strings"
)

func partTwo(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		res += computeHistoryTwo(parseLine(line))
	}
	return
}

func computeHistoryTwo(history []int) (res int) {
	if checkIfAllZeros(history) {
		return 0
	}
	nbThatICanTFindANameFuck := history[0]
	return nbThatICanTFindANameFuck - computeHistoryTwo(reductSlice(history))
}
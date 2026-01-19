package main

import (
	"strconv"
	"strings"
)

func partTwo(input string) int {
	bananaPerSequenceOverall := make(map[[4]int]int)
	for _, line := range strings.Split(input, "\n") {
		secret, _ := strconv.Atoi(line)
		computeBananasOfMonkey(secret, bananaPerSequenceOverall)
	}

	maxBananas := 0

	for _, bananas := range bananaPerSequenceOverall {
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}

	return maxBananas
}

func computeBananasOfMonkey(secret int, bananaPerSequenceOverall map[[4]int]int) {
	// première séquence
	var currSequence [4]int
	for i := 0; i < 4; i++ {
		newSecret := computeNextSecret(secret)
		currSequence[3-i] = (newSecret % 10) - (secret % 10)
		secret = newSecret
	}
	bananaPerSequenceOverall[currSequence] += secret % 10

	// visited pour ce singe
	visited := make(map[[4]int]bool)
	visited[currSequence] = true

	// le parcours de ce qui reste
	for i := 4; i < 2000; i++ {
		newSecret := computeNextSecret(secret)
		copy(currSequence[1:], currSequence[0:3])
		currSequence[0] = (newSecret % 10) - (secret % 10)
		if !visited[currSequence] {
			visited[currSequence] = true
			bananaPerSequenceOverall[currSequence] += newSecret % 10
		}
		secret = newSecret
	}
}
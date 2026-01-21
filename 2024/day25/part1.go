package main

import (
	"strings"
)

type Pattern [5]int

func partOne(input string) int {
	locks, keys := parseInput(input)
	return solvePartOne(locks, keys)
}

func solvePartOne(locks, keys []Pattern) (res int) {
	for _, lock := range locks {
		for _, key := range keys {
			isAPair := true
			for i := 0; i < 5; i++ {
				if lock[i] + key[i] > 5 {
					isAPair = false
					break
				}
			}
			if isAPair {
				res++
			}
		}
	}

	return
}

func parseInput(input string) (locks []Pattern, keys []Pattern) {
	locks = make([]Pattern, 0)
	keys = make([]Pattern, 0)

	patterns := strings.Split(input, "\n\n")

	for _, pattern := range patterns {
		isLock := false
		var currPattern Pattern
		lines := strings.Split(pattern, "\n")
		for r, line := range lines[:len(lines)-1] {
			if r == 0 {
				if line[0] == '#' {
					isLock = true
				}
			} else {
				for c, char := range line {
					if char == '#' {
						currPattern[c] += 1
					}
				}	
			}
		}
		if isLock {
			locks = append(locks, currPattern)
		} else {
			keys = append(keys, currPattern)
		}
	}

	return
}
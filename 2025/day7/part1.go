package main

import "strings"

func partOne(input string) (res int) {
	lines := strings.Split(input, "\n")
	currentBeams := make(map[int]bool)
	currentBeams[strings.Index(lines[0], "S")] = true

	for _, line := range lines[1:] {
		newBeams := make(map[int]bool)
		for c := range currentBeams {
			if line[c] == '^' {
				res++
				newBeams[c-1], newBeams[c+1] = true, true
			} else {
				newBeams[c] = true
			}
		}
		currentBeams = newBeams
	}

	return
}
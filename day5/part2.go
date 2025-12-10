package main

import (
	"strings"
)

func partTwo(input string) int {
	parts := strings.Split(input, "\n\n")
	ranges := parseRanges(strings.Split(parts[0], "\n"))
	res := 0

	for _, r := range ranges {
		res += r[1] - r[0] + 1
	}
	
	return res
}
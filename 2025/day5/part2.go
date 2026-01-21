package main

import (
	"strings"
)

func partTwo(input string) (res int) {
	parts := strings.Split(input, "\n\n")
	ranges := parseRanges(strings.Split(parts[0], "\n"))

	for _, r := range ranges {
		res += r[1] - r[0] + 1
	}
	
	return
}
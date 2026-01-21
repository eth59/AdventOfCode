package main

import (
	"slices"
	"strconv"
	"strings"
)

func partTwo(input string) (res int) {
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		nbMin, _ := strconv.Atoi(bounds[0])
		nbMax, _ := strconv.Atoi(bounds[1])
	
		for i := nbMin; i <= nbMax; i++ {
			currentStrNb := strconv.Itoa(i)
			for j := 2; j <= len(currentStrNb); j++ {
				compactParts := slices.Compact(splitStringN(currentStrNb, j))
				if len(compactParts) == 1 && strings.Count(currentStrNb, compactParts[0]) >= 2 {
					res += i
					break
				}
			}
		}
	}

	return 
}

func splitStringN(s string, n int) []string {
	lenS := len(s)
	if n <= 0 || n > lenS || lenS % n != 0 {
		return []string{}
	}
	lenPart := lenS / n
	parts := make([]string, n)
	for i := 0; i < n; i++ {
		parts[i] = s[i*lenPart : (i+1)*lenPart]
	}
	return parts
}
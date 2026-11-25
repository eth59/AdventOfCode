package main

import (
	"sort"
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	// parsing
	lines := strings.Split(input, "\n")
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))
	for _, line := range lines {
		nbs := strings.Split(line, "   ")
		nbLeft, _ := strconv.Atoi(nbs[0])
		nbRight, _ := strconv.Atoi(nbs[1])
		left = append(left, nbLeft)
		right = append(right, nbRight)
	}

	// on trie les slices
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	// calcul du res
	for i := range left {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = -distance
		}
		res += distance
	}
	return
}
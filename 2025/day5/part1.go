package main

import (
	"sort"
	"strconv"
	"strings"
)

func partOne(input string) int {
	parts := strings.Split(input, "\n\n")
	rangeLines := strings.Split(parts[0], "\n")
	idLines := strings.Split(parts[1], "\n")
	res := 0

	ranges := parseRanges(rangeLines)

	// on calcule le res
	for _, line := range idLines {
		id, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if isFresh(id, ranges) {
			res++
		}
	}

	return res
}

// regarde si un id donné est dans un range (après fusion)
func isFresh(id int, ranges [][2]int) bool {
	// on cherche le premier index d'un range qui commence après l'id à tester
	// on doit donc vérifier si l'id est dans le range juste avant celui qu'on vient de trouver
	rangeIndex := sort.Search(len(ranges), func(i int) bool {
		return ranges[i][0] > id
	}) - 1

	if rangeIndex >= 0 && id <= ranges[rangeIndex][1] {
		return true
	}
	return false
}
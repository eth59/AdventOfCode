package main

import (
	"sort"
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	parts := strings.Split(input, "\n\n")
	rangeLines := strings.Split(parts[0], "\n")
	idLines := strings.Split(parts[1], "\n")

	ranges := parseRanges(rangeLines)

	// on calcule le res
	for _, line := range idLines {
		id, _ := strconv.Atoi(line)
		if isFresh(id, ranges) {
			res++
		}
	}

	return
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

func parseRanges(rangeLines []string) (ranges [][2]int) {
	// on parse les ranges
	for _, line := range rangeLines {
		bounds := strings.Split(line, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, [2]int{start, end})
	}

	// on fusionne les ranges qui overlap
	ranges = mergeRanges(ranges)
	return
}

// fonction pour fusionner les ranges qui s'overlap
func mergeRanges(ranges [][2]int) (mergedRanges [][2]int) {
	sort.Slice(ranges, func(i,j int) bool {
		return ranges[i][0] <= ranges[j][0]
	})

	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		// si le range next commence avant ou en même temps que le current finit, on fusionne (on garde le max des deux fins)
		if next[0] <= current[1] + 1 {
			if next[1] > current[1] {
				current[1] = next[1]
			}
		} else {
			// ça chevauche pas, on save et on passe au suivant
			mergedRanges = append(mergedRanges, current)
			current = next
		}
	}
	mergedRanges = append(mergedRanges, current)
	return
}
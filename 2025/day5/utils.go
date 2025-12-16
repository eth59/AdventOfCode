package main

import (
	"sort"
	"strings"
	"strconv"
)

func parseRanges(rangeLines []string) [][2]int {
	// on parse les ranges
	var ranges [][2]int
	for _, line := range rangeLines {
		bounds := strings.Split(line, "-")
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, [2]int{start, end})
	}

	// on fusionne les ranges qui overlap
	ranges = mergeRanges(ranges)
	return ranges
}

// fonction pour fusionner les ranges qui s'overlap
func mergeRanges(ranges [][2]int) [][2]int {
	sort.Slice(ranges, func(i,j int) bool {
		return ranges[i][0] <= ranges[j][0]
	})

	var mergedRanges [][2]int
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
	return mergedRanges
}
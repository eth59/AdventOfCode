package main

import (
	"slices"
	"strconv"
	"strings"
)

type Map []Range

type Range struct {
	destStart, sourceStart, rangeLength int
}

func partOne(input string) int {
	seeds, maps := parse(input)
	nbSeeds := len(seeds)
	for _, m := range maps {
		currRes := make([]int, nbSeeds)
		for i, seed := range seeds {
			currRes[i] = convertWithMap(seed, m)
		}
		seeds = currRes
	}
	return slices.Min(seeds)
}

func convertWithMap(source int, mapToUse Map) int {
	for _, r := range mapToUse {
		if source >= r.sourceStart && source < r.sourceStart + r.rangeLength {
			return r.destStart + source - r.sourceStart
		}
	}
	return source
}

func parse(input string) (seeds []int, maps []Map) {
	parts := strings.Split(input, "\n\n")

	// seed
	seedsStr := strings.Fields(strings.Split(parts[0], ": ")[1])
	seeds = make([]int, len(seedsStr))
	for i, seedStr := range seedsStr {
		seed, _ := strconv.Atoi(seedStr)
		seeds[i] = seed
	}

	// maps
	maps = make([]Map, 7)
	for i, part := range parts[1:] {
		var currMap Map
		for _, line := range strings.Split(part, "\n")[1:] {
			nbsStr := strings.Fields(line)
			nbs := make([]int, 3)
			for j := 0; j < 3; j++ {
				n, _ := strconv.Atoi(nbsStr[j])
				nbs[j] = n
			}
			currMap = append(currMap, Range{nbs[0], nbs[1], nbs[2]})
		}
		maps[i] = currMap
	}
	return
}
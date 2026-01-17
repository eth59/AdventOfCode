package main

import (
	"slices"
	"strings"
)

func partTwo(input string) (res int) {
	towels, designs := parse(input)
	slices.SortFunc(towels, func(a, b string) int {
		return len(b) - len(a)
	})
	alreadyCounted := make(map[string]int)

	for _, design := range designs {
		res += countTowelsArrangements(towels, alreadyCounted, design)
	}

	return
}

func countTowelsArrangements(towels []string, alreadyCounted map[string]int, design string) (res int) {
	// cas d'arrêt
	if design == "" {
		return 1
	}

	// si déjà compté
	if count, ok := alreadyCounted[design] ; ok {
		return count
	}

	for _, towel := range towels {
		newDesign, found := strings.CutPrefix(design, towel)
		if found {
			count := countTowelsArrangements(towels, alreadyCounted, newDesign)
			alreadyCounted[newDesign] = count
			res += count
		}
	}

	alreadyCounted[design] = res

	return
}
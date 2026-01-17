package main

import (
	"slices"
	"strings"
)

func partOne(input string) (res int) {
	towels, designs := parse(input)
	slices.SortFunc(towels, func(a, b string) int {
		return len(b) - len(a)
	})

	for _, design := range designs {
		if isDesignPossible(towels, design) {
			res++
		}
	}

	return
}

func isDesignPossible(towels []string, design string) bool {
	// cas d'arrêt
	if design == "" {
		return true
	}

	for _, towel := range towels {
		newDesign, found := strings.CutPrefix(design, towel)
		if found && isDesignPossible(towels, newDesign) {
			return true
		}
	}

	return false
}

func parse(input string) (towels []string, designs []string) {
	parts := strings.Split(input, "\n\n")

	towels = strings.Split(parts[0], ", ")
	designs = strings.Split(parts[1], "\n")

	return
}
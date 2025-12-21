package main

import (
	"sort"
	"strconv"
	"strings"
)

func partTwo(input string) (res int) {
	parts := strings.Split(input, "\n\n")
	rules := make(map[int]map[int]bool)

	// parse des règles
	lines := strings.Split(parts[0], "\n")
	for _, line := range lines {
		nbs := strings.Split(line, "|")
		x, _ := strconv.Atoi(nbs[0])
		y, _ := strconv.Atoi(nbs[1])
		if rules[x] == nil {
			rules[x] = make(map[int]bool)
		}
		rules[x][y] = true
	}

	// parse updates & solve
	lines = strings.Split(parts[1], "\n")
	for _, line := range lines {
		updateStr := strings.Split(line, ",")
		updateLength := len(updateStr)
		update := make([]int, updateLength)
		for i, nb := range updateStr {
			update[i], _ = strconv.Atoi(nb)
		}

		if !isUpdateValid(update, rules, updateLength) {
			// les maps avec les règles permettent de directement faire la fonction pour trier les updates
			sort.Slice(update, func(i, j int) bool {
				return rules[update[i]][update[j]]
			})
			res += update[updateLength/2]
		}
	}
	return
}
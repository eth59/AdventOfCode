package main

import (
	"strconv"
	"strings"
)

func partOne(input string) int {
	return solve(input, 25)
}

func solve(input string, nbBlinks int) (res int) {
	stones := make(map[int]int)

	for _, s := range strings.Fields(input) {
		val, _ := strconv.Atoi(s)
		stones[val]++
	}

	for i := 0; i < nbBlinks; i++ {
		nextStones := make(map[int]int)

		for val, count := range stones {
			if val == 0 {
				nextStones[1] += count
			} else {
				valStr := strconv.Itoa(val)
				length := len(valStr)

				if length%2 == 0 {
					mid := length/2
					left, _ := strconv.Atoi(valStr[:mid])
					right, _ := strconv.Atoi(valStr[mid:])
					nextStones[left] += count
					nextStones[right] += count
				} else {
					nextStones[val*2024] += count
				}
			}
		}
		stones = nextStones
	}
	
	for _, count := range stones {
		res += count
	}
	return
}
package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		nbMin, _ := strconv.Atoi(bounds[0])
		nbMax, _ := strconv.Atoi(bounds[1])
		for i := nbMin; i <= nbMax; i++ {
			currentStrNb := strconv.Itoa(i)

			if len(currentStrNb) % 2 == 0 {
				mid := len(currentStrNb) / 2
				firstHalf := currentStrNb[:mid]
				secondHalf := currentStrNb[mid:]
				if firstHalf == secondHalf {
					res += i
				}
			}
		}
	}

	return
}
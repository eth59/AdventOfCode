package main

import (
	"strconv"
	"strings"
)

func partOne(input string) int {
	res := 0
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		nbMin, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		nbMax, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}
	
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

	return res
}
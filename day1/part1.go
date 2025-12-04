package main

import (
	"strconv"
	"strings"
)


func partOne(input string) int {
	current_position := 50
	res := 0

	for _, line := range strings.Split(input, "\n") {

		if len(line) == 0 {
			continue
		}

		nb, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		if line[0] == 'L' {
			current_position = (current_position - nb) % 100
		} else {
			current_position = (current_position + nb) % 100
		}

		if current_position == 0 {
			res++
		}
	}

	return res
}
package main

import (
	"strconv"
	"strings"
)


func partOne(input string) (res int) {
	current_position := 50

	for _, line := range strings.Split(input, "\n") {
		nb, _ := strconv.Atoi(line[1:])

		if line[0] == 'L' {
			current_position = (current_position - nb) % 100
		} else {
			current_position = (current_position + nb) % 100
		}

		if current_position == 0 {
			res++
		}
	}

	return
}
package main

import (
	"strconv"
	"strings"
)


func partTwo(input string) int {
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

		res += nb / 100
		nb -= (nb / 100) * 100

		if line[0] == 'L' {
			if current_position == 0 {
				current_position -= nb
			} else {
				current_position -= nb
				if current_position < 0 {
					res ++
				}
				current_position %= 100
			}
			
		} else {
			if current_position == 0 {
				current_position += nb
			} else {
				current_position += nb
				if current_position > 99 {
					res ++
				}
				current_position %= 100
			}
		}
		if current_position < 0 {
			current_position += 100
		}
	}

	return res
}
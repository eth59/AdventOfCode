package main

import (
	"strconv"
	"strings"
)


func partTwo(input string) int {
	current_position := 50
	res := 0
	already_counted := false

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

		if current_position == 0 {
			already_counted = true
		}

		if line[0] == 'L' {
			current_position -= nb
			if current_position <= 0 {
				if !already_counted {
					res++
				}
				if current_position < 0 {
					current_position += 100
				}
			}
		} else {
			current_position += nb
			if current_position >= 100 {
				res++
				current_position -= 100
			}
		}
		already_counted = false
	}

	return res
}
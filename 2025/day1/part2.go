package main

import (
	"strconv"
	"strings"
)


func partTwo(input string) (res int) {
	current_position := 50
	already_counted := false

	for _, line := range strings.Split(input, "\n") {
		nb, _ := strconv.Atoi(line[1:])

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

	return
}
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

		for i := 0; i < nb; i++ {
			if line[0] == 'L' {
				current_position -= 1
				if current_position == 0 {
					res++
				} else if current_position < 0 {
					current_position += 100
				}
			} else {
				current_position += 1
				if current_position == 100 {
					res++
					current_position = 0
				}
			}
 		}
	}

	return res
}
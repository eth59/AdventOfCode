package main

import "strings"

func partOne(input string) (res int) {
	instructions, network := parsePartOne(input)
	curr := "AAA"
	for {
		for _, instruction := range instructions {
			if instruction == 'L' {
				curr = network[curr][0]
			} else {
				curr = network[curr][1]
			}
			res++
			if curr == "ZZZ" {
				return
			}
		}
	}
}

func parsePartOne(input string) (string, map[string][2]string) {
	parts := strings.Split(input, "\n\n")

	network := make(map[string][2]string)
	for _, line := range strings.Split(parts[1], "\n") {
		network[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}

	return parts[0], network
}
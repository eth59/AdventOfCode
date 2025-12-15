package main

import "strings"

func parseInput(input string) (devices map[string][]string) {
	lines := strings.Split(input, "\n")
	devices = make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		devices[parts[0][:len(parts[0])-1]] = parts[1:]
	}

	return
}
package main

import "strings"

func partOne(input string) int {
	devices := parseInput(input)
	savedPaths := make(map[string]int)
	for name := range devices {
		savedPaths[name] = -1
	}
	return solverPartOne("you", devices, savedPaths)
}

func solverPartOne(name string, devices map[string][]string, savedPaths map[string]int) (res int) {
	for _, output := range devices[name] {
		if output == "out" {
			res++
		} else {
			if savedPaths[output] != -1 {
				res += savedPaths[output]
			} else {
				res += solverPartOne(output, devices, savedPaths)
			}
		}
	}
	savedPaths[name] = res
	return
}

func parseInput(input string) (devices map[string][]string) {
	lines := strings.Split(input, "\n")
	devices = make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		devices[parts[0][:len(parts[0])-1]] = parts[1:]
	}

	return
}
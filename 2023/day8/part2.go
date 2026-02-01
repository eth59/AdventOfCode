package main

import "strings"

func partTwo(input string) int {
	instructions, network, endWithA := parsePartTwo(input)
	n := len(endWithA)
	nbSteps := make([]int, n)
	for i := 0; i < n; i++ {
		nbSteps[i] = solvePartTwo(instructions, network, endWithA[i])
	}
	return ppcmSlice(nbSteps)
}

func solvePartTwo(instructions string, network map[string][2]string, curr string) (res int) {
	for {
		for _, instruction := range instructions {
			if instruction == 'L' {
				curr = network[curr][0]
			} else {
				curr = network[curr][1]
			}
			res++
			if curr[2] == 'Z' {
				return
			}
		}
	}
}

func pgcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func ppcmSlice(nbs []int) (res int) {
	res = nbs[0]
	for i := 1; i < len(nbs); i++ {
		res = (res / pgcd(res, nbs[i])) * nbs[i]
	}
	return
}

func parsePartTwo(input string) (string, map[string][2]string, []string) {
	endWithA := make([]string, 0)

	parts := strings.Split(input, "\n\n")

	network := make(map[string][2]string)
	for _, line := range strings.Split(parts[1], "\n") {
		key := line[0:3]
		if key[2] == 'A' {
			endWithA = append(endWithA, key)
		}
		network[key] = [2]string{line[7:10], line[12:15]}
	}

	return parts[0], network, endWithA
}
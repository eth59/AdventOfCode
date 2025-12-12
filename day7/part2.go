package main

import "strings"

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	calculatedPositions := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		calculatedPositions[i] = make([]int, len(lines[0]))
		for j := 0; j < len(lines[0]); j++ {
			calculatedPositions[i][j] = -1
		}
	}
	return countTimelines(lines, 0, strings.Index(lines[0], "S"), len(lines)-1, calculatedPositions)
}

func countTimelines(lines []string, r int, c int, end int, calculatedPositions [][]int) int {
	if calculatedPositions[r][c] != -1 {
		return calculatedPositions[r][c]
	} else if r == end {
		calculatedPositions[r][c] = 1
		return 1
	} else if lines[r][c] == '^' {
		calculatedPositions[r][c] = countTimelines(lines, r+1, c-1, end, calculatedPositions) + countTimelines(lines, r+1, c+1, end, calculatedPositions)
		return calculatedPositions[r][c]
	} else {
		calculatedPositions[r][c] = countTimelines(lines, r+1, c, end, calculatedPositions)
		return calculatedPositions[r][c]
	}
}
package main

import (
	"strconv"
	"strings"
	"unicode"
)

func partTwo(input string) (res int) {
	lines := strings.Split(input, "\n")
	height, width := len(lines), len(lines[0])
	for r, line := range lines {
		for c, char := range line {
			if char == '*' {
				res += getGearRatio(lines, height, width, r, c)
			}
		}
	}
	return
}

func getGearRatio(lines []string, height, width, r, c int) (res int) {
	// la clef c {ligne, startCol} & la valeur c le nb
	adjacentNumbers := make(map[[2]int]int)
	for i := 0; i < 8; i++ {
		nr, nc := r + dr[i], c + dc[i]
		if nr >= 0 && nr < height && nc >= 0 && nc < width && unicode.IsDigit(rune(lines[nr][nc])) {
			val, startCol := getNumber(lines, width, nr, nc)
			adjacentNumbers[[2]int{nr, startCol}] = val
		}
	}
	if len(adjacentNumbers) == 2 {
		res = 1
		for _, n := range adjacentNumbers {
			res *= n
		}
	}
	return
}

func getNumber(lines []string, width, r, c int) (int, int) {
	row := lines[r]

	start := c
	for start >= 0 && unicode.IsDigit(rune(row[start-1])) {
		start--
	}

	end := c
	for end < width && unicode.IsDigit(rune(row[end])) {
		end++
	}

	val, _ := strconv.Atoi(row[start:end])
	return val, start
}
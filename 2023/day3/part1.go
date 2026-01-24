package main

import (
	"strings"
	"unicode"
)

// de en haut à en haut à gauche
var dr = []int{-1, -1, 0, 1, 1, 1, 0, -1}
var dc = []int{0, 1, 1, 1, 0, -1, -1, -1}

func partOne(input string) (res int) {
	lines := strings.Split(input, "\n")
	height, width := len(lines), len(lines[0])
	for r, line := range lines {
		var currentNumber int
		var isPartNumber bool
		for c, char := range line {
			if unicode.IsDigit(char) {
				currentNumber = currentNumber*10 + int(char-'0')
				if !isPartNumber && checkNeighbors(lines, height, width, r, c) {
					isPartNumber = true
				}
			} else {
				if isPartNumber {
					res += currentNumber
				}
				currentNumber = 0
				isPartNumber = false
			}
		}
		if isPartNumber {
			res += currentNumber
		}
		currentNumber = 0
		isPartNumber = false
	}
	return
}

func checkNeighbors(lines []string, height, width, r, c int) bool {
	for i := 0; i < 8; i++ {
		nr, nc := r + dr[i], c + dc[i]
		if nr >= 0 && nr < height && nc >= 0 && nc < width && lines[nr][nc] != '.' && !unicode.IsDigit(rune(lines[nr][nc])) {
			return true
		}
	}
	return false
}
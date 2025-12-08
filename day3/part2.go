package main

import (
	"strconv"
	"strings"
)

func findLargestJoltage(line string) string {
	const requiredLength = 12
	if len(line) < requiredLength {
		return ""
	}

	result := make([]rune, 0, requiredLength)
	inputRunes := []rune(line)
	startIndex := 0
	
	for i := 0; i < requiredLength; i++ {
		endIndex := len(inputRunes) - requiredLength + i + 1
		maxDigit := inputRunes[startIndex]
		maxIndex := startIndex

		for j := startIndex+1; j < endIndex; j++ {
			if inputRunes[j] > maxDigit {
				maxDigit = inputRunes[j]
				maxIndex = j
			}
		}

		result = append(result, maxDigit)
		startIndex = maxIndex + 1
	}
	return string(result)
}

func partTwo(input string) int {
	res := 0

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		maxJoltageStr := findLargestJoltage(line)
		maxJoltage, err := strconv.Atoi(maxJoltageStr)
		if err != nil {
			continue // Probablement la dernière ligne vide
		}
		res += maxJoltage
	}
	return res
}
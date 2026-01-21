package main

import (
	"strconv"
	"strings"
)

func partTwo(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		maxJoltageStr := findLargestJoltage(line)
		maxJoltage, _ := strconv.Atoi(maxJoltageStr)
		res += maxJoltage
	}
	return
}

func findLargestJoltage(line string) string {
	const requiredLength = 12
	
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
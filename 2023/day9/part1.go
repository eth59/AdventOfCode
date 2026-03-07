package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		res += computeHistory(parseLine(line))
	}
	return
}

func parseLine(line string) (res []int) {
	lineSlice := strings.Fields(line)
	res = make([]int, 0, len(lineSlice))
	for _, x := range lineSlice {
		xInt, _ := strconv.Atoi(x)
		res = append(res, xInt)
	}
	return
}

func computeHistory(history []int) (res int) {
	if checkIfAllZeros(history) {
		return 0
	}
	toAdd := history[len(history)-1]
	return toAdd + computeHistory(reductSlice(history))
}

func reductSlice(s []int) []int {
	n := len(s) - 1
	for i := 0 ; i < n ; i++ {
		s[i] = s[i+1] - s[i]
	}
	s = s[:n]
	return s
}

func checkIfAllZeros(s []int) bool {
	for _, x := range s {
		if x != 0 {
			return false
		}
	}
	return true
}
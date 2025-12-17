package main

import (
	"slices"
	"strconv"
	"strings"
)

func partTwo(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		length := len(fields)
		report := make([]int, 0, len(fields))
		for _, nbStr := range fields {
			nb, _ := strconv.Atoi(nbStr)
			report = append(report, nb)
		}
		if isSafe(report, 0, length, true) {
			res++
		} else { // même code que partie 1 mais en + on check en enlevant un élément
			for i := 0; i < len(report); i++ {
				newReport := slices.Concat(report[:i], report[i+1:])
				if isSafe(newReport, 0, length-1, true) {
					res++
					break
				}
			}
		}
	}
	return
}
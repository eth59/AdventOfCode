package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
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
		}
	}
	return
}

// vérifie si un rapport est safe ou pas
func isSafe(report []int, index, length int, ascending bool) bool {
	// condition d'arrêt, on est au dernier élément
	if index == length - 1 {
		return true
	} else if report[index] == report[index+1] { // égalité => unsafe
		return false
	} else if absInt(report[index] - report[index+1]) < 1 || absInt(report[index] - report[index+1]) > 3 { // différence trop grande
		return false
	} else if index == 0 { // premier élément, on détermine le sens du tri du report
		if report[index] < report[index+1] {
			return isSafe(report, index+1, length, true)
		} else {
			return isSafe(report, index+1, length, false)
		}
	} else if ascending { // croissant
		if report[index+1] < report[index] {
			return false
		} else {
			return isSafe(report, index+1, length, ascending)
		}
	} else { // décroissant
		if report[index+1] > report[index] {
			return false
		} else {
			return isSafe(report, index+1, length, ascending)
		}
	}
}

// abs pour int
func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
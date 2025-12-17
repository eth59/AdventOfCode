package main

import (
	"regexp"
	"strconv"
)

func partTwo(input string) (res int) {
	// regex pour trouver les multiplications valides et les do/dont
	reMult := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	mults := reMult.FindAllStringSubmatch(input, -1)
	enable := true
	for _, mult := range mults {
		switch mult[0] {
		case "don't()":
			enable = false
		case "do()":
			enable = true
		default:
			if enable {
				x, _ := strconv.Atoi(mult[1])
				y, _ := strconv.Atoi(mult[2])
				res += x*y
			}
		}
	}
	return
}
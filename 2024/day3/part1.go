package main

import (
	"regexp"
	"strconv"
)

func partOne(input string) (res int) {
	// regex pour trouver les multiplications valides
	reMult := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	mults := reMult.FindAllStringSubmatch(input, -1)
	for _, mult := range mults {
		x, _ := strconv.Atoi(mult[1])
		y, _ := strconv.Atoi(mult[2])
		res += x*y
	}
	return
}
package main

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	r, c int
}

type MemoKey struct {
	curr, next, depth int
}
var memo = make(map[MemoKey]int)

var mapNumericKeypad = map[int]Point{
	7 : {0,0}, 8 : {0,1}, 9 : {0,2},
	4 : {1,0}, 5 : {1,1}, 6 : {1,2},
	1 : {2,0}, 2 : {2,1}, 3 : {2,2},
	-1: {3,0}, 0 : {3,1},10 : {3,2},
}

var mapDirectionalKeypad = map[int]Point{
	-1: {0,0}, 3 : {0,1}, 4 : {0,2},
	2 : {1,0}, 1 : {1,1}, 0 : {1,2},
}

func partOne(input string) int {
	return solve(input, 3)
}

func solve(input string, depth int) (res int) {
	lines := strings.Split(input, "\n")
	for _, code := range lines {
		nbInCodeStr := ""
		var pathLength, digit int
		curr := 10
		for _, digitRune := range code {
			if digitRune != 'A' {
				digit, _ = strconv.Atoi(string(digitRune))
				nbInCodeStr += string(digitRune)
			} else {
				digit = 10
			}
			pathLength += findMinCost(curr, digit, depth, true)
			curr = digit
		}
		nbInCode, _ := strconv.Atoi(nbInCodeStr)
		res += nbInCode * pathLength
	}
	return
}

func findMinCost(curr, next, depth int, isNumeric bool) int {
	// cas d'arrêt
	if depth == 0 {
		return 1
	}

	memoKey := MemoKey{curr, next, depth}
	if cost, alreadyComputed := memo[memoKey] ; alreadyComputed {
		return cost
	}

	var currPoint, nextPoint, gap Point
	if isNumeric {
		currPoint = mapNumericKeypad[curr]
		nextPoint = mapNumericKeypad[next]
		gap = mapNumericKeypad[-1]
	} else {
		currPoint = mapDirectionalKeypad[curr]
		nextPoint = mapDirectionalKeypad[next]
		gap = mapDirectionalKeypad[-1]
	}

	dr := nextPoint.r - currPoint.r
	dc := nextPoint.c - currPoint.c

	minCost := math.MaxInt

	// horizontal puis vertical
	if !(gap.r == currPoint.r && gap.c == nextPoint.c) {
		path := ""
		if dc < 0 { path += strings.Repeat("<", -dc) }
		if dc > 0 { path += strings.Repeat(">", dc) }
		if dr < 0 { path += strings.Repeat("^", -dr) }
		if dr > 0 { path += strings.Repeat("v", dr) }
		path += "A"

		cost := 0
		currOnNextKeypad := 4 // A
		var arrowInt int
		for _, arrow := range path {
			switch arrow {
			case '<':
				arrowInt = 2
			case '>':
				arrowInt = 0
			case '^':
				arrowInt = 3
			case 'v':
				arrowInt = 1
			case 'A':
				arrowInt = 4
			}
			cost += findMinCost(currOnNextKeypad, arrowInt, depth-1, false)
			currOnNextKeypad = arrowInt
		}

		if cost < minCost {
			minCost = cost
		}
	}

	// vertical puis horizontal
	if !(gap.c == currPoint.c && gap.r == nextPoint.r) {
		path := ""
		if dr < 0 { path += strings.Repeat("^", -dr) }
		if dr > 0 { path += strings.Repeat("v", dr) }
		if dc < 0 { path += strings.Repeat("<", -dc) }
		if dc > 0 { path += strings.Repeat(">", dc) }
		path += "A"

		cost := 0
		currOnNextKeypad := 4 // A
		var arrowInt int
		for _, arrow := range path {
			switch arrow {
			case '<':
				arrowInt = 2
			case '>':
				arrowInt = 0
			case '^':
				arrowInt = 3
			case 'v':
				arrowInt = 1
			case 'A':
				arrowInt = 4
			}
			cost += findMinCost(currOnNextKeypad, arrowInt, depth-1, false)
			currOnNextKeypad = arrowInt
		}

		if cost < minCost {
			minCost = cost
		}
	}
	
	memo[memoKey] = minCost
	return minCost
}
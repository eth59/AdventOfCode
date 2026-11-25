package main

import (
	"slices"
	"strconv"
	"strings"
)

func partTwo(input string) (res int) {
	var valCard = map[rune]int {
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}
	hands := parsePartTwo(input, valCard)
	slices.SortFunc(hands, func(a, b Hand) int {
		return a.score - b.score
	})
	for i, hand := range hands {
		res += hand.bid * (i+1)
	}
	return
}

func parsePartTwo(input string, valCard map[rune]int) (hands []Hand) {
	lines := strings.Split(input, "\n")
	hands = make([]Hand, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")

		var handScore int
		var counts [13]int
		for _, card := range parts[0] {
			val := valCard[card]
			handScore = (handScore << 4) + val
			counts[val]++
		}
		
		var maxCount, secondMax int
		for _, c := range counts[1:] { // on prend pas les jokers
			if c > maxCount {
				secondMax = maxCount
				maxCount = c
			} else if c > secondMax {
				secondMax = c
			}
		}

		maxCount += counts[0] // on met les jokers dans le max

		if maxCount == 5 {
			// five of a kind
			handScore += 60000000000
		} else if maxCount == 4 {
			// four of a kind
			handScore += 50000000000
		} else if maxCount == 3 && secondMax == 2 {
			// full house
			handScore += 40000000000
		} else if maxCount == 3 {
			// three of a kind
			handScore += 30000000000
		} else if maxCount == 2 && secondMax == 2 {
			// two pairs
			handScore += 20000000000
		} else if maxCount == 2 {
			// one pair
			handScore += 10000000000
		}

		bid, _ := strconv.Atoi(parts[1])
		hands[i] = Hand{handScore, bid}
	}
	return
}
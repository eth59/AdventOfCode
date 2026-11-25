package main

import (
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	score int
	bid int
}

func partOne(input string) (res int) {
	var valCard = map[rune]int {
		'A': 12,
		'K': 11,
		'Q': 10,
		'J': 9,
		'T': 8,
		'9': 7,
		'8': 6,
		'7': 5,
		'6': 4,
		'5': 3,
		'4': 2,
		'3': 1,
		'2': 0,
	}
	hands := parsePartOne(input, valCard)
	slices.SortFunc(hands, func(a, b Hand) int {
		return a.score - b.score
	})
	for i, hand := range hands {
		res += hand.bid * (i+1)
	}
	return
}

func parsePartOne(input string, valCard map[rune]int) (hands []Hand) {
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
		for _, c := range counts {
			if c > maxCount {
				secondMax = maxCount
				maxCount = c
			} else if c > secondMax {
				secondMax = c
			}
		}

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
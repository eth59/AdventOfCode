package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		secret, _ := strconv.Atoi(line)
		for i := 0; i < 2000; i++ {
			secret = computeNextSecret(secret)
		}
		res += secret
	}
	return
}

func computeNextSecret(currSecret int) int {
	currSecret = prune(mix(currSecret*64, currSecret))
	currSecret = prune(mix(currSecret/32, currSecret))
	return prune(mix(currSecret*2048, currSecret))
}

func mix(a, secret int) int {
	return a ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}
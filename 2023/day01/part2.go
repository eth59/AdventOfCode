package main

import (
	"strings"
	"unicode"
)

var digitWords = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func partTwo(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		var firstDigit, lastDigit int
		for i, char := range line {
			if unicode.IsDigit(char) {
				if firstDigit == 0 {
					firstDigit = int(char - '0')
				}
				lastDigit = int(char - '0')
			} else {
				for digit, digitWord := range digitWords {
					if strings.HasPrefix(line[i:], digitWord) {
						if firstDigit == 0 {
							firstDigit = digit
						}
						lastDigit = digit
					}
				}
			}
		}
		res += firstDigit * 10 + lastDigit
	}
	return
}

// version plus lente à base de regex en dessous

// func partTwo(input string) (res int) {
// 	reFirstDigit := regexp.MustCompile(`^.*?(\d|zero|one|two|three|four|five|six|seven|eight|nine)`)
// 	reLastDigit := regexp.MustCompile(`^.*(\d|zero|one|two|three|four|five|six|seven|eight|nine)`)

// 	for _, line := range strings.Split(input, "\n") {
// 		matchesFirstDigit := reFirstDigit.FindAllStringSubmatch(line, -1)[0]
// 		matchesLastDigit := reLastDigit.FindAllStringSubmatch(line, -1)[0]
		
// 		firstDigit, err := strconv.Atoi(matchesFirstDigit[1])
// 		if err != nil {
// 			firstDigit = lettersToDigit[matchesFirstDigit[1]]
// 		}

// 		lastDigit, err := strconv.Atoi(matchesLastDigit[1])
// 		if err != nil {
// 			lastDigit = lettersToDigit[matchesLastDigit[1]]
// 		}

// 		res += firstDigit * 10 + lastDigit
// 	}

// 	return
// }

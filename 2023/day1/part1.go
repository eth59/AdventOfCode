package main

import (
	"strconv"
	"strings"
	"unicode"
)

func partOne(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		var firstDigit, lastDigit rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				if firstDigit == 0 {
					firstDigit = char
				}
				lastDigit = char
			}
		}
		nb, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		res += nb
	}
	return
}


// version plus lente à base de regex en dessous

// func partOne(input string) (res int) {
// 	re := regexp.MustCompile(`^.*?(\d).*?(\d)?\D*$`)
// 	for _, line := range strings.Split(input, "\n") {
// 		matches := re.FindAllStringSubmatch(line, -1)[0]
// 		if matches[2] == "" {
// 			nb, _ := strconv.Atoi(matches[1] + matches[1])
// 			res += nb
// 		} else {
// 			nb, _ := strconv.Atoi(matches[1] + matches[2])
// 			res += nb
// 		}
// 	}
// 	return
// }
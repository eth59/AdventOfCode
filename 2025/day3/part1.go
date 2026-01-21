package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
    for _, line := range strings.Split(input, "\n") {        
        runes := []rune(line)
        max1 := runes[0]
		max2 := runes[1]

        for i, nb := range runes[2:] {
			if i+3 == len(runes) && nb > max2 {
				max2 = nb
			} else if nb > max1 {
                max1 = nb
				max2 = runes[i+1]
            } else if nb > max2 {
				max2 = nb
			}
        }

		nbMax, _ := strconv.Atoi(string(max1) + string(max2))
		res += nbMax
    }
    return
}
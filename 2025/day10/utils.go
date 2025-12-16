package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Lights []int
	Buttons [][]int
	Joltages []int
}

func parseInput(input string) (res []Machine) {
	lines := strings.Split(input, "\n")
	res = make([]Machine, 0, len(lines))
	reLights := regexp.MustCompile(`\[([.#]+)\]`)
	reButtons := regexp.MustCompile(`\(([\d,]+)\)`)
	reJoltages := regexp.MustCompile(`\{([\d,]+)\}`)
	for _, line := range lines {
		// parse lumières
		lightsStr := reLights.FindStringSubmatch(line)[1]
		lights := make([]int, 0, len(lightsStr))
		for _, value := range lightsStr {
			if value == '#' {
				lights = append(lights, 1)
			} else {
				lights = append(lights, 0)
			}
		} 

		// parse buttons
		matches := reButtons.FindAllSubmatch([]byte(line), -1)
		buttons := make([][]int, 0, len(matches))
		for _, m := range matches {
			buttonsValues := strings.Split(string(m[1]), ",")
			var btn []int
			for _, v := range buttonsValues {
				val, _ := strconv.Atoi(v)
				btn = append(btn, val)
			}
			buttons = append(buttons, btn)
		}

		// parse joltages
		joltagesStr := strings.Split(reJoltages.FindStringSubmatch(line)[1], ",")
		joltages := make([]int, 0, len(joltagesStr))
		for _, value := range joltagesStr {
			nb, _ := strconv.Atoi(value)
			joltages = append(joltages, nb)
		}

		res = append(res, Machine{lights, buttons, joltages})
	}
	return
}
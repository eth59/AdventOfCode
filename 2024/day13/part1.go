package main

import (
	"strconv"
	"strings"
)

func partOne(input string) (res int) {
	machines := strings.Split(input, "\n\n")
	for _, machine := range machines {
		m := strings.Split(machine, "\n")
		
		// button A
		a := strings.Split(strings.Split(m[0], ": ")[1], ", ")
		xa, _ := strconv.Atoi(a[0][2:])
		ya, _ := strconv.Atoi(a[1][2:])

		// button B
		b := strings.Split(strings.Split(m[1], ": ")[1], ", ")
		xb, _ := strconv.Atoi(b[0][2:])
		yb, _ := strconv.Atoi(b[1][2:])

		// prize
		p := strings.Split(strings.Split(m[2], ": ")[1], ", ")
		px, _ := strconv.Atoi(p[0][2:])
		py, _ := strconv.Atoi(p[1][2:])

		// calcul de la machine
		numerator := px*yb-py*xb
		denominator := xa*yb-ya*xb
		res_a := perfectDivision(numerator, denominator)
		numerator = px*ya-py*xa
		denominator = xb*ya-yb*xa
		res_b := perfectDivision(numerator, denominator)
		if res_a != 0 && res_b != 0 {
			res += 3*res_a + res_b
		}
	}
	return
}

// renvoie le résultat de la division seulement s'il n'y a pas de reste
func perfectDivision(numerator, denominator int) int {
	if numerator % denominator == 0 {
		return numerator / denominator
	}
	return 0
}
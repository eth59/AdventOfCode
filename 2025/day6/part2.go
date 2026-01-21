package main

import (
	"strconv"
	"strings"
)

func partTwo(input string) (res int) {
	lines := strings.Split(input, "\n")
	height := len(lines)

	operands := make([]string, 10)
	op := ""
	currentOperand := 0
	
	// on parcourt l'input par colonne
	for c := 0; c < len(lines[0]); c++ {
		emptyColumn := true // on va traquer les colonnes vides qui séparent les opérations

		// on ajoute cette colonne comme une opérande pour notre calcule si elle n'est pas vide
		for r := 0; r < height-1; r++ {
			if lines[r][c] != ' ' {
				operands[currentOperand] += string(lines[r][c])
				emptyColumn = false
			}
		}

		// on va passer à la prochaine opérande au prochain tour
		currentOperand++

		// si on a pas d'opération, c'est qu'on est à la première colonne de ce calcul et on récupére l'opération
		if op == "" {
			op = string(lines[height-1][c])
		}

		// si la colonne est vide, on a finit un calcul, on le compute et on réinitialise les variables
		if emptyColumn {
			res += computeColumn(operands, op)
			operands = make([]string, 10)
			op = ""
			currentOperand = 0
		}
	}

	// on oublie pas le dernier calcul
	res += computeColumn(operands, op)
	return
}


// fonction pour calculer le résultat d'une opération
func computeColumn(operands []string, op string) (res int) {
	res, _ = strconv.Atoi(operands[0])
	for _, operand := range operands[1:] {
		if operand == "" { break }
		nb, _ := strconv.Atoi(operand)
		if op == "+" {
			res += nb
		} else {
			res *= nb
		}
	}
	
	return
}
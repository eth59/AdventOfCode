package main

import "math"

func partTwo(input string) int {
	instructions, _, _, _ := parse(input)
	return findLowestA(instructions, 0, len(instructions) - 1)
}

func findLowestA(instructions []int, currentA, depth int) int {
	// cas de base
	if depth < 0 {
		return currentA
	}

	// on décale pour ajouter les 3 nouveaux bits à tester
	currentA <<= 3

	// on teste toutes les combinaisons des 3 bits
	for i := 0; i < 8; i++ {
		candidateA := currentA + i
		output := runOneCycle(instructions, candidateA)
		if output == instructions[depth] {
			res := findLowestA(instructions, candidateA, depth-1)
			if res != -1 {
				return res
			}
		}
	}

	return -1 // pas de solution trouvée
}

func runOneCycle(instructions []int, registerA int) int {
	var registerB, registerC, instructionPointer int

	getComboOperand := func(operand int) int {
		switch operand {
		case 4:
			return registerA
		case 5:
			return registerB
		case 6:
			return registerC
		default:
			return operand
		}
	}
	
	adv := func(operand int) { // opcode 0
		registerA /= int(math.Pow(2, float64(getComboOperand(operand))))
	}
	
	bxl := func(operand int) { // opcode 1
		registerB ^= operand
	}
	
	bst := func(operand int) { // opcode 2
		registerB = getComboOperand(operand) % 8
	}
	
	bxc := func() { // opcode 4
		registerB ^= registerC
	}
	
	bdv := func(operand int) { // opcode 6
		registerB = registerA / int(math.Pow(2, float64(getComboOperand(operand))))
	}
	
	cdv := func(operand int) { // opcode 7
		registerC = registerA / int(math.Pow(2, float64(getComboOperand(operand))))
	}

	for instructionPointer < len(instructions) {
		operand := instructions[instructionPointer+1]
		switch instructions[instructionPointer] {
		case 0:
			adv(operand)
			instructionPointer += 2
		case 1:
			bxl(operand)
			instructionPointer += 2
		case 2:
			bst(operand)
			instructionPointer += 2
		case 3:
			// on fait rien, on veut pas boucler mais juste faire un tour
		case 4:
			bxc()
			instructionPointer += 2
		case 5:
			// plutôt qu'appeler out, on return la valeur
			return getComboOperand(operand) % 8
		case 6:
			bdv(operand)
			instructionPointer += 2
		case 7:
			cdv(operand)
			instructionPointer += 2
		}
	}
	return -1 // impossible logiquement
}
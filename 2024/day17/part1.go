package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func partOne(input string) string {
	instructions, registerA, registerB, registerC := parse(input)
	return runProgram(instructions, registerA, registerB, registerC)
}

func parse(input string) (instructions []int, registerA, registerB, registerC int) {
	parts := strings.Split(input, "\n\n")

	// registers
	registers := strings.Split(parts[0], "\n")
	registerA, _ = strconv.Atoi(strings.Split(registers[0], ": ")[1])
	registerB, _ = strconv.Atoi(strings.Split(registers[1], ": ")[1])
	registerC, _ = strconv.Atoi(strings.Split(registers[2], ": ")[1])

	// instructions
	instructionsStr := strings.Split(strings.Split(parts[1], ": ")[1], ",")
	instructions = make([]int, 0, len(instructionsStr))
	for _, instructionStr := range instructionsStr {
		instruction, _ := strconv.Atoi(instructionStr)
		instructions = append(instructions, instruction)
	} 
	return
}

func runProgram(instructions []int, registerA, registerB, registerC int) string {
	var instructionPointer int
	var output string

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
	
	jnz := func(operand int) bool { // opcode 3
		if registerA == 0 {
			return false
		} else {
			instructionPointer = operand
			return true
		}
	}
	
	bxc := func() { // opcode 4
		registerB ^= registerC
	}
	
	out := func(operand int) { // opcode 5
		output += fmt.Sprintf("%d,", getComboOperand(operand)%8)	
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
			if !jnz(operand) {
				instructionPointer += 2
			}
		case 4:
			bxc()
			instructionPointer += 2
		case 5:
			out(operand)
			instructionPointer += 2
		case 6:
			bdv(operand)
			instructionPointer += 2
		case 7:
			cdv(operand)
			instructionPointer += 2
		}
	}
	return output[:len(output)-1]
}
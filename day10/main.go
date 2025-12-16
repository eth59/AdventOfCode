package main

import (
	_ "embed"
	"fmt"
)

//go:embed input_test.txt
var inputTest string

//go:embed input.txt
var input string

func main() {
	fmt.Println("PART 1:")

	machineInputTest := parseInput(inputTest)
	machineInput := parseInput(input)

	resultTest := partOne(machineInputTest)
	fmt.Println("Result test:", resultTest)

	result := partOne(machineInput)
	fmt.Println("Result:", result)

	fmt.Println("PART 2:")

	resultTest2 := partTwo(machineInputTest)
	fmt.Println("Result test:", resultTest2)

	result = partTwo(machineInput)
	fmt.Println("Result:", result)
}
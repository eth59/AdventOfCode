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

	resultTest := partOne(inputTest)
	fmt.Println("Result test:", resultTest)

	result := partOne(input)
	fmt.Println("Result:", result)

	fmt.Println("PART 2:")

	resultTest = partTwo(inputTest)
	fmt.Println("Result test:", resultTest)

	result = partTwo(input)
	fmt.Println("Result:", result)
}
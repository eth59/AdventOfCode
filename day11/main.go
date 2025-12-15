package main

import (
	_ "embed"
	"fmt"
)

//go:embed input_test.txt
var inputTest string

//go:embed input_test2.txt
var inputTest2 string

//go:embed input.txt
var input string

func main() {
	fmt.Println("PART 1:")

	devicesTest := parseInput(inputTest)
	devicesTest2 := parseInput(inputTest2)
	devices := parseInput(input)

	resultTest := partOne(devicesTest)
	fmt.Println("Result test:", resultTest)

	result := partOne(devices)
	fmt.Println("Result:", result)

	fmt.Println("PART 2:")

	resultTest2 := partTwo(devicesTest2)
	fmt.Println("Result test:", resultTest2)

	result = partTwo(devices)
	fmt.Println("Result:", result)
}
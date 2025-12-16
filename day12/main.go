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

	giftsTest, regionsTest := parseInput(inputTest)
	gifts, regions := parseInput(input)

	resultTest := partOne(giftsTest, regionsTest)
	fmt.Println("Result test:", resultTest)

	result := partOne(gifts, regions)
	fmt.Println("Result:", result)
}
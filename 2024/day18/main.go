package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input_test.txt
var inputTest string

//go:embed input.txt
var input string

func main() {
	args := os.Args[1:]
	
	// valeurs par défaut
	runPart1 := true
	runPart2 := true
	isTest := false

	// gestion des arguments
	for _, arg := range args {
		lowArg := strings.ToLower(arg)
		switch {
		case lowArg == "test" || lowArg == "t":
			isTest = true
		case lowArg == "1":
			runPart1 = true
			runPart2 = false
		case lowArg == "2":
			runPart2 = true
			runPart1 = false
		}
	}

	// selection de l'input
	if isTest {
		fmt.Println("🧪🧪🧪  MODE TEST 🧪🧪🧪")
	}

	// exécution des parties
	if runPart1 {
		if isTest {
			execute(1, inputTest, 12, 7, 7)
		} else {
			execute(1, input, 1024, 71, 71)
		}
	}
	if runPart2 {
		if isTest {
			execute(2, inputTest, 12, 7, 7)
		} else {
			execute(2, input, 1024, 71, 71)
		}
	}
}

// exécution avec affichage des résultats
func execute(part int, input string, nbBytesCorrupted, width, height int) {
	fmt.Printf("--- YEAR 2024 - DAY 18 - PART %d ---\n", part)
	var result interface{}
	if part == 1 {
		result = partOne(input, nbBytesCorrupted, width, height)
	} else {
		result = partTwo(input, width, height)
	}
	fmt.Printf("Result: %v\n", result)
	fmt.Println()
}
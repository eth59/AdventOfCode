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
			execute(1, inputTest)
		} else {
			execute(1, input)
		}
	}
	if runPart2 {
		if isTest {
			execute(2, inputTest)
		} else {
			execute(2, input)
		}
	}
}

// exécution avec affichage des résultats
func execute(part int, input string) {
	fmt.Printf("--- YEAR 2023 - DAY 9 - PART %d ---\n", part)
	var result interface{}
	if part == 1 {
		result = partOne(input)
	} else {
		result = partTwo(input)
	}
	fmt.Printf("Result: %v\n", result)
	fmt.Println()
}
package main

import "fmt"

func partOne(machinesInput []Machine) (res int) {
	for _, m := range machinesInput {
		res += solveMachine(m) 
	}
	return
}

// pivot de gauss (L*B)(B*1) = (L*1)
// L : lights & B : buttons
func solveMachine(m Machine) int {
	nbLights := len(m.Lights)
	nbButtons := len(m.Buttons)

	// on construit la matrice augmentée [L*B|L*1]
	// init
	matrix := make([][]int, nbLights)
	for i := range matrix {
		matrix[i] = make([]int, nbButtons+1)
	}
	// remplissage
	for btnIndex, btnValues := range m.Buttons {
		for _, lightIndex := range btnValues {
			matrix[lightIndex][btnIndex] = 1
		}
	}
	for i, val := range m.Lights {
		matrix[i][nbButtons] = val
	}

	fmt.Println(matrix)
	
	return 0
}
package main

import (
	"math"
)

func partOne(machinesInput []Machine) (res int) {
	for _, m := range machinesInput {
		res += solveMachine(m) 
	}
	return
}

// pivot de gauss (L*B)(B*1) = (L*1) mod 2
// L : lights & B : buttons
func solveMachine(m Machine) (res int) {
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

	// pivot de gauss
	pivotRow := 0
	pivotCols := make([]int, nbLights) // pour stocker quelle colonne est pivot de chaque ligne
	for i := range pivotCols {
		pivotCols[i] = -1
	}
	isPivotCol := make([]bool, nbButtons) // pour retrouver les variables libres après

	for col := 0; col < nbButtons && pivotRow < nbLights; col++ {
		// on cherche une ligne avec un 1 dans cette colonne
		selected := -1
		for row := pivotRow; row < nbLights; row++ {
			if matrix[row][col] == 1 {
				selected = row
				break
			}
		}

		if selected == -1 {
			continue // pas de pivot dans cette colonne, variable libre
		}

		// swap des lignes
		matrix[pivotRow], matrix[selected] = matrix[selected], matrix[pivotRow]
		pivotCols[pivotRow] = col
		isPivotCol[col] = true

		// annuler les autres lignes
		for row := 0; row < nbLights; row++ {
			if row != pivotRow && matrix[row][col] == 1 {
				for k := col; k <= nbButtons; k++ {
					matrix[row][k] ^= matrix[pivotRow][k] // xor parce qu'on est mod 2
				}
			}
		}

		// la ligne pivot augmente
		pivotRow++
	}

	// on identifie les variables libres
	var freeVars []int
	for col := 0; col < nbButtons; col++ {
		if !isPivotCol[col] {
			freeVars = append(freeVars, col)
		}
	}

	// recherche solution minimale
	res = math.MaxInt32

	// on va tester toutes les combinaisons de variables libres
	for i := 0; i < 1 << len(freeVars); i++ {
		solution := make([]int, nbButtons)
		currentNbPressed := 0

		// variables libres prennent la valeur des bits de i
		for j, v := range freeVars {
			if (i >> j)&1 == 1 {
				solution[v] = 1
				currentNbPressed++
			} else {
				solution[v] = 0
			}
		}

		// clacul variables pivots
		for row := 0; row < pivotRow; row++ {
			pCol := pivotCols[row]
			val := matrix[row][nbButtons] // valeur cible

			// xor des variables libres activées sur cette ligne
			for _, fCol := range freeVars {
				// si c une variable libre activée avec un 1 dans cette ligne
				if matrix[row][fCol] == 1 && solution[fCol] == 1 {
					val ^= 1
				}
			}

			solution[pCol] = val
			if val == 1 {
				currentNbPressed++
			}
		}

		if currentNbPressed < res {
			res = currentNbPressed
		}
	}
	
	return
}
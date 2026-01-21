package main

import (
	"math"
)

func partTwo(input string) (res int) {
	machinesInput := parseInput(input)
	for _, m := range machinesInput {
		res += solveMachinePart2(m)
	}
	return
}

func solveMachinePart2(m Machine) (res int) {
	nbCounters := len(m.Joltages)
	nbButtons := len(m.Buttons)

	// on construit la matrice augmentée [L*B|L*1]
	// init
	matrix := make([][]float64, nbCounters)
	for i := range matrix {
		matrix[i] = make([]float64, nbButtons+1)
	}
	// remplissage
	for btnIndex, btnValues := range m.Buttons {
		for _, counterIndex := range btnValues {
			matrix[counterIndex][btnIndex] = 1.0
		}
	}
	for i, val := range m.Joltages {
		matrix[i][nbButtons] = float64(val)
	}

	// pivot de gauss
	pivotRow := 0
	pivotCols := make([]int, nbCounters) // pour stocker quelle colonne est pivot de chaque ligne
	for i := range pivotCols {
		pivotCols[i] = -1
	}
	isPivotCol := make([]bool, nbButtons) // pour retrouver les variables libres après

	for col := 0; col < nbButtons && pivotRow < nbCounters; col++ {
		// on cherche le meilleur pivot tel qu'il soit minimal et strictement positif
		selected := -1
		maxVal := 0.0
		for row := pivotRow; row < nbCounters; row++ {
			absVal := math.Abs(matrix[row][col])
			if absVal > maxVal {
				maxVal = absVal
				selected = row
			}
		}

		if selected == -1 {
			continue // pas de pivot dans cette colonne, variable libre
		}

		// swap des lignes
		matrix[pivotRow], matrix[selected] = matrix[selected], matrix[pivotRow]
		pivotCols[pivotRow] = col
		isPivotCol[col] = true

		// normalisation
		div := matrix[pivotRow][col]
		for k := col; k <= nbButtons; k++ {
			matrix[pivotRow][k] /= div
		}

		// annuler les autres lignes
		for row := 0; row < nbCounters; row++ {
			if row != pivotRow && matrix[row][col] != 0 {
				factor := matrix[row][col]
				for k := col; k <= nbButtons; k++ {
					matrix[row][k] -= factor * matrix[pivotRow][k]
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
	res = -1

	// valeur max pour borner la recherche
	maxTarget := 0
	for _, v := range m.Joltages {
		if v > maxTarget {
			maxTarget = v
		}
	}

	// on teste les combinaisons de variables libres avec une fonction récursive
	var search func(idx int, currentSol []int)
	search = func(idx int, currentSol []int) {
		if idx == len(freeVars) {
			// on a fixé les var libres, on calcule les var pivots
			nbButtonsPressed := 0
			valid := true
			
			// on copie la solution partielle
			fullSol := make([]float64, nbButtons)
			for i, fIdx := range freeVars {
				fullSol[fIdx] = float64(currentSol[i])
				nbButtonsPressed += currentSol[i]
			}

			// calcul des var pivots
			for row := 0; row < nbCounters; row++ {
				pCol := pivotCols[row]
				if pCol == -1 {
					continue
				}

				val := matrix[row][nbButtons] // valeur cible
				
				// on retire l'influence des var libres
				for _, fCol := range freeVars {
					if matrix[row][fCol] != 0 {
						val -= matrix[row][fCol] * fullSol[fCol]
					}
				}
				
				// on vérifie que c un entier positif
				const epsilon = 1e-9 // pour les erreurs de calcul avec les flottants
				if val < -epsilon {
					valid = false
					break
				}
				rounded := math.Round(val)
				if math.Abs(val - rounded) > epsilon {
					valid = false
					break
				}
				
				fullSol[pCol] = rounded
				nbButtonsPressed += int(rounded)
			}

			if valid {
				if res == -1 || nbButtonsPressed < res {
					res = nbButtonsPressed
				}
			}
			return
		}

		// on boucle sur la var libre actuelle
		// on limite la recherche
		for v := 0; v <= maxTarget; v++ {
			currentSol[idx] = v
			search(idx+1, currentSol)
		}
	}

	// on lance la recherche
	initialSol := make([]int, len(freeVars))
	search(0, initialSol)

	return
}
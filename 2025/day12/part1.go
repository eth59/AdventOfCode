package main

import (
	"slices"
	"strconv"
	"strings"
)

type Gift struct {
	shape [9]bool
	Area int
}

func (g Gift) set(x, y int, value bool) Gift {
	g.shape[x*3+y] = value
	return g
}

func (g Gift) get(x, y int) bool {
	return g.shape[x*3+y]
}

// String pour le debug
func (g Gift) String() (res string) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if g.get(x, y) {
				res += "#"
			} else {
				res += "."
			}
		}
		if x < 2 {
			res += "\n"
		}
	}
	return
}

// rotate dans le sens horaire
func (g Gift) rotate() (res Gift) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			res = res.set(x, y, g.get(2-y, x))
		}
	}
	return
}

// flip horizontal
func (g Gift) flip() (res Gift) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			res = res.set(x, y, g.get(x, 2-y))
		}
	}
	return
}

type Region struct {
	SizeX, SizeY int
	NeededShapes []int
	NeededGifts int
}

func partOne(input string) (res int) {
	gifts, regions := parseInput(input)

	// on génère les variantes des cadeaux
	giftsVariants := make([][]Gift, 0, len(gifts))
	for _, gift := range gifts {
		giftVariants := make([]Gift, 0, 8)
		giftVariants = append(giftVariants, gift)
		for i := 0; i < 3; i++ {
			gift = gift.rotate()
			if !slices.Contains(giftVariants, gift) {
				giftVariants = append(giftVariants, gift)
			}
		}
		gift = gift.flip()
		if !slices.Contains(giftVariants, gift) {
				giftVariants = append(giftVariants, gift)
			}
		for i := 0; i < 3; i++ {
			gift = gift.rotate()
			if !slices.Contains(giftVariants, gift) {
				giftVariants = append(giftVariants, gift)
			}
		}
		giftsVariants = append(giftsVariants, giftVariants)
	}

	// on lance le solver
	for _, region := range regions {
		// on vérifie qu'il y a mathématiquement la place de faire rentrer les cadeaux
		totalGiftArea := 0
		for giftIdx, count := range region.NeededShapes {
			totalGiftArea += count * gifts[giftIdx].Area
		}
		if totalGiftArea > region.SizeX * region.SizeY {
			continue // impossible
		}

		// on crée une liste avec tous les index des cadeaux à ajouter dans la région
		// on retourve x fois un index s'il faut le mettre x fois
		giftsToTry := make([]int, 0, region.NeededGifts)
		for i, nbNeeded := range region.NeededShapes {
			for j := 0; j < nbNeeded; j++ {
				giftsToTry = append(giftsToTry, i)
			}
		}

		// création de la région qu'on va remplir
		currentState := make([][]bool, 0, region.SizeX)
		for i := 0; i < region.SizeX; i++ {
			currentState = append(currentState, make([]bool, region.SizeY))
		}

		if solverPartOne(giftsVariants, giftsToTry, currentState, 0) {
			res++
		}
	}

	return
}

// backtracking
func solverPartOne(variants [][]Gift, giftsToTry []int, currentState [][]bool, index int) bool {
	// condition d'arrêt
	if index == len(giftsToTry) {
		return true
	}

	// on essaye toutes les variantes du prochain cadeau à essayer
	for _, variant := range variants[giftsToTry[index]] {
		for i := 0; i < len(currentState); i++ {
			for j := 0; j < len(currentState[0]); j++ {
				if canPlace(variant, currentState, i, j) {
					// on place le cadeau
					place(variant, currentState, i, j, true)

					// récursion
					if solverPartOne(variants, giftsToTry, currentState, index+1) {
						return true
					}

					// on retire le cadeau
					place(variant, currentState, i, j, false)
				}
			}
		}
	}

	return false
}

// teste si l'ajout d'un cadeau est possible
func canPlace(gift Gift, currentState [][]bool, x, y int) bool {
	if len(currentState) - x < 3 || len(currentState[0]) - y < 3 {
		return false
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gift.get(i, j) {
				if currentState[x+i][y+j] {
					return false
				}
			}
		}
	}
	return true
}

// ajoute ou retire un cadeau dans la grille
func place(gift Gift, currentState [][]bool, x, y int, val bool) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gift.get(i, j) {
				currentState[x+i][y+j] = val
			}
		}
	}
}

func parseInput(input string) (gifts []Gift, regions []Region) {
	parts := strings.Split(input, "\n\n")

	// parse gifts
	for _, part := range parts[:len(parts)-1] {
		gift := Gift{}
		area := 0
		lines := strings.Split(part, "\n")
		for i, line := range lines[1:] {
			for j, c := range line {
				if c == '#' {
					area++
					gift = gift.set(i, j, true)
				} else {
					gift = gift.set(i, j, false)
				}
			}
		}
		gift.Area = area
		gifts = append(gifts, gift)
	}

	// parse regions
	for _, line := range strings.Split(parts[len(parts)-1], "\n") {
		parts := strings.Split(line, ": ")
		size := strings.Split(parts[0], "x")
		sizeX, _ := strconv.Atoi(size[0])
		sizeY, _ := strconv.Atoi(size[1])
		neededShapesStr := strings.Split(parts[1], " ")
		neededShapes := make([]int, 0, len(neededShapesStr))
		var neededGifts int
		for _, nbStr := range neededShapesStr {
			nb, _ := strconv.Atoi(nbStr)
			neededShapes = append(neededShapes, nb)
			neededGifts += nb
		}

		regions = append(regions, Region{sizeX, sizeY, neededShapes, neededGifts})
	}

	return
}
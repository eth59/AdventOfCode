package main

import (
	"maps"
	"slices"
	"strings"
)

func partTwo(input string) string {
	adjacencyMatrix := createAdjacencyMatrix(input)
	maxClique := bronKerbosch(adjacencyMatrix, nil, slices.Collect(maps.Keys(adjacencyMatrix)), nil)
	slices.Sort(maxClique)
	return strings.Join(maxClique, ",")
}

func bronKerbosch(adjacencyMatrix map[string][]string, R, P, X []string) (maxClique []string) {
	if len(P) == 0 && len(X) == 0 {
		return R
	}

	candidates := make([]string, len(P))
	copy(candidates, P)
	for _, v := range candidates {
		newR := append([]string{}, R...)
		newR = append(newR, v)
		
		newP := inter(P, adjacencyMatrix[v])
		newX := inter(X, adjacencyMatrix[v])
		newMaxClique := bronKerbosch(adjacencyMatrix, newR, newP, newX)
		if len(newMaxClique) > len(maxClique) {
			maxClique = newMaxClique
		}
		P = remove(P, v)
		X = append(X, v)
	}
	return
}

func remove(slice []string, val string) []string {
	newSlice := make([]string, 0)
	for _, x := range slice {
		if x != val {
			newSlice = append(newSlice, x)
		}
	}
	return newSlice
}

func inter(E []string, F []string) []string {
	res := make([]string, 0)
	for _, e := range E {
		if slices.Contains(F, e) {
			res = append(res, e)
		}
	}
	return res
}
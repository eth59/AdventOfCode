package main

import (
	"maps"
	"slices"
	"strings"
)

func partOne(input string) int {
	adjacencyMatrix := createAdjacencyMatrix(input)
	sortedVertex := slices.Sorted(maps.Keys(adjacencyMatrix))
	return findThreeCliques(sortedVertex, adjacencyMatrix)
}

func createAdjacencyMatrix(input string) map[string][]string {
	adjacencyMatrix := make(map[string][]string)
	for _, conn := range strings.Split(input, "\n") {
		parts := strings.Split(conn, "-")
		a, b := parts[0], parts[1]
		adjacencyMatrix[a] = append(adjacencyMatrix[a], b)
		adjacencyMatrix[b] = append(adjacencyMatrix[b], a)
	}
	return adjacencyMatrix
}

func findThreeCliques(sortedVertex []string, adjacencyMatrix map[string][]string) (res int) {
	nbVertex := len(sortedVertex)
	for i, u := range sortedVertex {
		for j := i+1; j < nbVertex; j++ {
			v := sortedVertex[j]
			if slices.Contains(adjacencyMatrix[u], v) {
				for k := j+1; k < nbVertex; k++ {
					w := sortedVertex[k]
					if slices.Contains(adjacencyMatrix[u], w) && slices.Contains(adjacencyMatrix[v], w) {
						// on a trouvé une 3-clique
						if strings.HasPrefix(u, "t") || strings.HasPrefix(v, "t") || strings.HasPrefix(w, "t") {
							// faut un t
							res++
						}
					}
				}
			}
		}
	}
	return
}
package main

import (
	"container/heap"
	"math"
	"strings"
)

func partTwo(input string) (res int) {
	// parsing
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	entryR, entryC, outputR, outputC := parse(lines)

	// dijkstra	
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startState := State{entryR, entryC, 0}
	startItem := Item{startState, 0, 0}
	heap.Push(&pq, &startItem)

	dist := make(map[State]int) // stockage du coût minimal actuel trouvé
	dist[startState] = 0

	parents := make(map[State][]State) // stockage des états précédents pr atteindre un état

	minEndCost := math.MaxInt

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)

		// si on dépasse déjà le coût minimal, inutile de continuer
		if curr.cost > minEndCost {
			continue
		}

		// si on a déjà trouvé un chemin moins cher, on skip
		if prevCost, exists := dist[curr.state]; exists && prevCost < curr.cost {
			continue
		}

		// si on est à la fin, on peut modifier le coût minimal
		if curr.state.r == outputR && curr.state.c == outputC {
			if curr.cost < minEndCost {
				minEndCost = curr.cost
			}
			continue
		}

		// avancer
		nr, nc := curr.state.r + dr[curr.state.dir], curr.state.c + dc[curr.state.dir]
		if nr >= 0 && nr < height && nc >= 0 && nc < width && lines[nr][nc] != '#' {
			newCost := curr.cost + 1
			nextState := State{nr, nc, curr.state.dir}
			existingCost, visited := dist[nextState]
			if !visited || newCost < existingCost {
				dist[nextState] = newCost
				parents[nextState] = []State{curr.state}
				heap.Push(&pq, &Item{nextState, newCost, 0})
			} else if newCost == existingCost {
				parents[nextState] = append(parents[nextState], curr.state)
			}
		}

		// tourner
		rotations := []int{(curr.state.dir + 1) % 4, (curr.state.dir + 3) % 4}
		for _, newDir := range rotations {
			newCost := curr.cost + 1000
			nextState := State{curr.state.r, curr.state.c, newDir}
			existingCost, visited := dist[nextState]
			if !visited || newCost < existingCost {
				dist[nextState] = newCost
				parents[nextState] = []State{curr.state}
				heap.Push(&pq, &Item{nextState, newCost, 0})
			} else if newCost == existingCost {
				parents[nextState] = append(parents[nextState], curr.state)
			}
		}
	}

	// get res
	uniqueTiles := make(map[[2]int]bool)
	seenStates := make(map[State]bool)
	queue := make([]State, 0)

	for dir := 0; dir < 4; dir++ {
		endState := State{outputR, outputC, dir}
		if cost, visited := dist[endState]; visited && cost == minEndCost {
			queue = append(queue, endState)
			seenStates[endState] = true
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		uniqueTiles[[2]int{curr.r, curr.c}] = true

		for _, parent := range parents[curr] {
			if !seenStates[parent] {
				seenStates[parent] = true
				queue = append(queue, parent)
			}
		}
	}

	return len(uniqueTiles)
}
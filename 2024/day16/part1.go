package main

import (
	"container/heap"
	"strings"
)

// 0 : droite, 1 : bas, 2 : gauche, 3 : haut
var dr = []int{0, 1, 0, -1}
var dc = []int{1, 0, -1, 0}

type State struct {
	r, c int
	dir int
}

type Item struct {
	state State
	cost int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func partOne(input string) int {
	// parsing
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	entryR, entryC, outputR, outputC := parse(lines)

	// solve	
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startState := State{entryR, entryC, 0}
	startItem := Item{startState, 0, 0}
	heap.Push(&pq, &startItem)

	dist := make(map[State]int) // stockage du coût minimal actuel trouvé
	dist[startState] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)

		// si on est à la fin, on retourne le coût
		if curr.state.r == outputR && curr.state.c == outputC {
			return curr.cost
		}

		// si on a déjà trouvé un chemin moins cher, on skip
		if prevCost, exists := dist[curr.state]; exists && prevCost < curr.cost {
			continue
		}

		// avancer
		nr, nc := curr.state.r + dr[curr.state.dir], curr.state.c + dc[curr.state.dir]
		if nr >= 0 && nr < height && nc >= 0 && nc < width && lines[nr][nc] != '#' {
			newCost := curr.cost + 1
			nextState := State{nr, nc, curr.state.dir}
			if existingCost, visited := dist[nextState]; !visited || newCost < existingCost {
				dist[nextState] = newCost
				heap.Push(&pq, &Item{nextState, newCost, 0})
			}
		}

		// tourner
		rotations := []int{(curr.state.dir + 1) % 4, (curr.state.dir + 3) % 4}
		for _, newDir := range rotations {
			newCost := curr.cost + 1000
			nextState := State{curr.state.r, curr.state.c, newDir}
			if existingCost, visited := dist[nextState]; !visited || newCost < existingCost {
				dist[nextState] = newCost
				heap.Push(&pq, &Item{nextState, newCost, 0})
			}
		}
	}

	return -1 // normalement c impossible
}

func parse(lines []string) (entryR, entryC, outputR, outputC int) {
	for r, line := range lines {
		for c, cell := range line {
			switch cell {
			case 'S':
				entryR, entryC = r, c
			case 'E':
				outputR, outputC = r, c
			}
		}
	}
	return
}
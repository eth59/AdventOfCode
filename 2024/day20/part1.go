package main

import "strings"

type Point struct {
	r, c int
}

// 0 : droite, 1 : bas, 2 : gauche, 3 : haut
var dr = []int{0, 1, 0, -1}
var dc = []int{1, 0, -1, 0}

var CHEAT_MIN_PICOSECONDS_FOR_RES = 100

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	startingPoint := findStart(lines)
	path, trackTimes := mapTrackIter(lines, startingPoint)
	return findCheatsIter(lines, path, trackTimes)
}

func mapTrackIter(lines []string, startingPoint Point) ([]Point, map[Point]int) {
	path := []Point{startingPoint}
	trackTimes := make(map[Point]int)
	trackTimes[startingPoint] = 0
	nextIndex := 1

	curr := startingPoint
	for lines[curr.r][curr.c] != 'E' {
		for i := 0; i < 4; i++ {
			nr, nc := curr.r + dr[i], curr.c + dc[i]
			nPoint := Point{nr, nc}
			if _, visited := trackTimes[nPoint] ; !visited && lines[nr][nc] != '#' {
				trackTimes[nPoint] = nextIndex
				path = append(path, nPoint)
				nextIndex++
				curr = nPoint
				break
			}
		}
	}

	return path, trackTimes
}

func findCheatsIter(lines []string, path []Point, trackTimes map[Point]int) (res int) {
	for currentTime, currentPoint := range path {
		// on regarde dans les 4 directions s'il y a un mur
		for i := 0; i < 4; i++ {
			nr, nc := currentPoint.r + dr[i], currentPoint.c + dc[i]
			if lines[nr][nc] != '#' {
				continue // on veut pas chercher de raccourci ici, c'est pas un mur
			}

			// on regarde les 3 voisins du mur, si c'est une piste on regarde si c'est un bon raccourci
			for j := 0; j < 4; j++ {
				if j == (i + 2) % 4 {
					continue // c'est le retour sur la case currentPoint
				}
				endShortcutR, endShortcutC := nr + dr[j], nc + dc[j]
				if endTime, isTrack := trackTimes[Point{endShortcutR, endShortcutC}] ; isTrack {
					if endTime - currentTime - 2 >= CHEAT_MIN_PICOSECONDS_FOR_RES {
						res++
					}
				}
			}
		}
	}
	return
}

func findStart(lines []string) Point {
	for r, line := range lines {
		for c, cell := range line {
			if cell == 'S' {
				return Point{r, c}
			}
		}
	}
	return Point{-1, -1} // normalement impossible
}

// Vieille version récursive en dessous, plus lent qu'en itératif

// func partOne(input string) int {
// 	lines := strings.Split(input, "\n")
// 	startingPoint := findStart(lines)

// 	trackTimes := make(map[Point]int)
// 	trackTimes[startingPoint] = 0
// 	mapTrack(lines, trackTimes, startingPoint, 1)

// 	visited := make(map[Point]bool)
// 	visited[startingPoint] = true
// 	return findCheats(lines, trackTimes, visited, startingPoint)
// }

// func findCheats(lines []string, trackTimes map[Point]int, visited map[Point]bool, currentPoint Point) (res int) {
// 	// cas d'arrêt
// 	if lines[currentPoint.r][currentPoint.c] == 'E' {
// 		return 0
// 	}

// 	// on regarde dans les 4 directions s'il y a un mur
// 	for i := 0; i < 4; i++ {
// 		nr, nc := currentPoint.r + dr[i], currentPoint.c + dc[i]
// 		nPoint := Point{nr, nc}
// 		if lines[nr][nc] != '#' {
// 			if !visited[nPoint] {
// 				visited[nPoint] = true
// 				res += findCheats(lines, trackTimes, visited, nPoint)
// 			}
// 			continue // on veut pas chercher de raccourci ici, c'est pas un mur
// 		}

// 		// on regarde les 3 voisins du mur, si c'est une piste on regarde si c'est un bon raccourci
// 		for j := 0; j < 4; j++ {
// 			if j == (i + 2) % 4 {
// 				continue // c'est le retour sur la case currentPoint
// 			}
// 			endShortcutR, endShortcutC := nr + dr[j], nc + dc[j]
// 			if endTime, isTrack := trackTimes[Point{endShortcutR, endShortcutC}] ; isTrack {
// 				if endTime - trackTimes[currentPoint] - 2 >= CHEAT_MIN_PICOSECONDS_FOR_RES {
// 					res++
// 				}
// 			}
// 		}
// 	}

// 	return
// }

// func mapTrack(lines []string, trackTimes map[Point]int, currentPoint Point, currentTime int) {
// 	// cas d'arrêt
// 	if lines[currentPoint.r][currentPoint.c] == 'E' {
// 		return
// 	}

// 	// on regarde dans les 4 directions la suite de la piste
// 	for i := 0; i < 4; i++ {
// 		nr, nc := currentPoint.r + dr[i], currentPoint.c + dc[i]
// 		nPoint := Point{nr, nc}
// 		if _, ok := trackTimes[nPoint] ; !ok && lines[nr][nc] != '#' {
// 			trackTimes[nPoint] = currentTime
// 			mapTrack(lines, trackTimes, nPoint, currentTime+1)
// 			return
// 		}
// 	}
// }
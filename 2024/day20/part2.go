package main

import "strings"

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dist(a, b Point) int {
	return absInt(a.r - b.r) + absInt(a.c - b.c)
}

var CHEAT_MAX_LAST_FOR = 20

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	startingPoint := findStart(lines)
	path, _ := mapTrackIter(lines, startingPoint)
	return findCheatsTwoIter(path)
}

func findCheatsTwoIter(path []Point) (res int) {
	for currentTime, currentPoint := range path {
		for endShortcutTime := currentTime + 1 ; endShortcutTime < len(path) ; endShortcutTime++ {
			endShortcutPoint := path[endShortcutTime]
			cheatLastFor := dist(currentPoint, endShortcutPoint)
			if cheatLastFor <= CHEAT_MAX_LAST_FOR && endShortcutTime - currentTime - cheatLastFor >= CHEAT_MIN_PICOSECONDS_FOR_RES {
				res++
			}
		}
	}
	return
}


// Vieille version récursive en dessous, plus lent qu'en itératif

// func partTwo(input string) int {
// 	lines := strings.Split(input, "\n")
// 	startingPoint := findStart(lines)

// 	trackTimes := make(map[Point]int)
// 	trackTimes[startingPoint] = 0
// 	mapTrack(lines, trackTimes, startingPoint, 1)

// 	visited := make(map[Point]bool)
// 	visited[startingPoint] = true
// 	return findCheatsTwo(lines, trackTimes, visited, startingPoint)
// }

// func findCheatsTwo(lines []string, trackTimes map[Point]int, visited map[Point]bool, currentPoint Point) (res int) {
// 	// cas d'arrêt
// 	if lines[currentPoint.r][currentPoint.c] == 'E' {
// 		return 0
// 	}

// 	// on va chercher tous les points de track à une distance inf à CHEAT_MAX_LAST_FOR
// 	// on peut se limiter aux points non visités sinon on retourne en arrière
// 	// et on vérifie si on y gagne plus de temps que CHEAT_MIN_PICOSECONDS_FOR_RES
// 	for endShortcut, endShortcutTime := range trackTimes {
// 		if visited[endShortcut] {
// 			continue
// 		}
// 		cheatLastFor := dist(currentPoint, endShortcut)
// 		if cheatLastFor <= CHEAT_MAX_LAST_FOR && endShortcutTime - trackTimes[currentPoint] - cheatLastFor >= CHEAT_MIN_PICOSECONDS_FOR_RES {
// 			res++
// 		}
// 	}

// 	// récursivité
// 	for i := 0; i < 4; i++ {
// 		nr, nc := currentPoint.r + dr[i], currentPoint.c + dc[i]
// 		nPoint := Point{nr, nc}
// 		if lines[nr][nc] != '#' && !visited[nPoint] {
// 			visited[nPoint] = true
// 			res += findCheatsTwo(lines, trackTimes, visited, nPoint)
// 			return
// 		}
// 	}

// 	return
// }
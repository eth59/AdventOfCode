package main

import (
	"strings"
)

func partTwo(input string) (res int) {
	parts := strings.Split(input, "\n\n")
	roomStr := strings.Split(parts[0], "\n")
	height := len(roomStr)
	width := len(roomStr[0])
	var robotX, robotY int

	// parse & find robot
	robotX, robotY, room := parsePartTwo(roomStr, height, width)

	// move robot
	moves := strings.Split(parts[1], "\n")
	var dir int
	for _, movesLine := range moves {
		for _, move := range movesLine {
			switch move {
			case '^':
				dir = 0
			case '>':
				dir = 1
			case 'v':
				dir = 2
			case '<':
				dir = 3
			}
			if canRobotMove(room, robotX, robotY, dir) {
				robotX, robotY = moveRobotTwo(room, robotX, robotY, dir, 0)
			}
		}
	}

	// calcul res
	for r, line := range room {
		for c, val := range line {
			if val == 1 {
				res += r*100+c
			}
		}
	}

	return
}

func parsePartTwo(roomStr []string, height, width int) (robotX int, robotY int, room [][]int) {
	room = make([][]int, 0, height) // 0 : rien/robot, 1 : boîte gauche, 2 : mur, 3 : boîte droite
	for y, line := range roomStr {
		row := make([]int, 0, width*2)
		for x, c := range line {
			switch c {
			case '@':
				robotX, robotY = x*2, y
				row = append(append(row, 0), 0)
			case '.':
				row = append(append(row, 0), 0)
			case 'O':
				row = append(append(row, 1), 3)
			default:
				row = append(append(row, 2), 2)
			}
		}
		room = append(room, row)
	}
	return 
}

func canRobotMove(room [][]int, x, y, dir int) bool {
	switch dir {
	case 0:
		// haut
		y--
	case 1:
		// droite
		x++
	case 2:
		// bas
		y++
	case 3:
		// gauche
		x--
	}
	switch room[y][x] {
	case 0:
		return true
	case 1:
		switch dir {
		case 0, 2:
			return canRobotMove(room, x, y, dir) && canRobotMove(room, x+1, y, dir) // faut tester les deux cases de la boîte
		case 1:
			return canRobotMove(room, x+1, y, dir) // on peut skip direct la partie droite de la boîte
		case 3:
			return canRobotMove(room, x, y, dir)
		}
	case 2:
		return false
	case 3:
		switch dir {
		case 0, 2:
			return canRobotMove(room, x, y, dir) && canRobotMove(room, x-1, y, dir) // faut tester les deux cases de la boîte
		case 1:
			return canRobotMove(room, x, y, dir)
		case 3:
			return canRobotMove(room, x-1, y, dir) // on peut skip direct la partie gauche de la boîte
		}
	}
	return false // on est pas censé arriver là
}

func moveRobotTwo(room [][]int, x, y, dir, val int) (int, int) {
	switch dir {
	case 0:
		// haut
		y--
	case 1:
		// droite
		x++
	case 2:
		// bas
		y++
	case 3:
		// gauche
		x--
	}
	switch room[y][x] {
	case 0:
		room[y][x] = val
	case 1:
		if dir == 0 || dir == 2 {
			moveRobotTwo(room, x, y, dir, 1) 
			moveRobotTwo(room, x+1, y, dir, 3)
			room[y][x] = val
			room[y][x+1] = 0
		} else {
			moveRobotTwo(room, x, y, dir, 1)
			room[y][x] = val
		}
	case 3:
		if dir == 0 || dir == 2 {
			moveRobotTwo(room, x, y, dir, 3) 
			moveRobotTwo(room, x-1, y, dir, 1)
			room[y][x] = val
			room[y][x-1] = 0
		} else {
			moveRobotTwo(room, x, y, dir, 3)
			room[y][x] = val
		}
	}
	return x, y
}
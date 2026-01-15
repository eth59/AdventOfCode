package main

import (
	"errors"
	"strings"
)

func partOne(input string) (res int) {
	parts := strings.Split(input, "\n\n")
	roomStr := strings.Split(parts[0], "\n")
	height := len(roomStr)
	width := len(roomStr[0])
	var robotX, robotY int

	// parse & find robot
	robotX, robotY, room := parsePartOne(roomStr, height, width)

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
			x, y, err := moveRobot(room, robotX, robotY, dir, 0)
			if err == nil {
				robotX, robotY = x, y
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

func parsePartOne(roomStr []string, height, width int) (robotX int, robotY int, room [][]int) {
	room = make([][]int, 0, height) // 0 : rien/robot, 1 : boîte, 2 : mur
	for y, line := range roomStr {
		row := make([]int, 0, width)
		for x, c := range line {
			switch c {
			case '@':
				robotX, robotY = x, y
				row = append(row, 0)
			case '.':
				row = append(row, 0)
			case 'O':
				row = append(row, 1)
			default:
				row = append(row, 2)
			}
		}
		room = append(room, row)
	}
	return 
}

func moveRobot(room [][]int, x, y, dir, val int) (int, int, error) {
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
		return x, y, nil
	case 1:
		_, _, err := moveRobot(room, x, y, dir, 1)
		if err != nil {
			return x, y, err
		}
		room[y][x] = val
		return x, y, nil
	default:
		return x, y, errors.New("déplacement impossible")
	}
}
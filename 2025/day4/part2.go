package main

func partTwo(input string) (res int) {
	grid := parseGrid(input)

	height := len(grid)
	width := len(grid[0])

	for {
		toRemove := make([][2]int, 0)

		for r := 0; r < height; r++ {
			for c := 0; c < width; c++ {
				if grid[r][c] == '@' {
					nbNeighbors := 0

					for _, dir := range directions {
						nr, nc := r+dir[0], c+dir[1]
						if nr >= 0 && nr < height && nc >= 0 && nc < width && grid[nr][nc] == '@' {
							nbNeighbors++
						}
					}

					if nbNeighbors < 4 {
						toRemove = append(toRemove, [2]int{r, c})
					}
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		res += len(toRemove)

		for _, coord := range toRemove {
			grid[coord[0]][coord[1]] = '.'
		}
	}

	return
}
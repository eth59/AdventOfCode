package main

func partOne(input [][]rune) int {
	res := 0

	height := len(input)
	width := len(input[0])

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, 		   {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if input[r][c] == '@' {
				nbNeighbors := 0
				for _, dir := range directions {
					nr, nc := r+dir[0], c+dir[1]
					if nr >= 0 && nr < height && nc >= 0 && nc < width && input[nr][nc] == '@' {
						nbNeighbors++
					}
				}
				if nbNeighbors < 4 {
					res++
				}
			}
		}
	}
	return res
}
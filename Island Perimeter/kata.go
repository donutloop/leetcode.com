package Island_Perimeter

func islandPerimeter(grid [][]int) int {
	var sum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == 0 {
				continue
			}

			left := j - 1

			if left >= 0 && grid[i][left] == 0 {
				sum++
			} else if left < 0 {
				sum++
			}

			right := j + 1
			if right < len(grid[i]) && grid[i][right] == 0 {
				sum++
			} else if right == len(grid[i]) {
				sum++
			}

			up := i - 1
			if up >= 0 && grid[up][j] == 0 {
				sum++
			} else if up < 0 {
				sum++
			}

			down := i + 1
			if down < len(grid) && grid[down][j] == 0 {
				sum++
			} else if down == len(grid) {
				sum++
			}
		}
	}

	return sum
}

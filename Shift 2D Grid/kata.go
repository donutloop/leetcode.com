package Shift_2D_Grid

func shiftGrid(grid [][]int, k int) [][]int {

	swap := -1001
	for ; k >= 1; k-- {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				c := swap
				if c != -1001 {
					if i == (len(grid)-1) && j == (len(grid[i])-1) {
						grid[0][0] = swap
						swap = -1001
					} else if i+1 < len(grid) && j == (len(grid[i])-1) {
						oldSwap := swap
						swap = grid[i+1][0]
						grid[i+1][0] = oldSwap
					} else if j < len(grid[i]) {
						oldSwap := swap
						swap = grid[i][j+1]
						grid[i][j+1] = oldSwap
					}
				} else {
					if i == (len(grid)-1) && j == (len(grid[i])-1) {
						return grid
					} else if i+1 < len(grid) && j == (len(grid[i])-1) {
						swap = grid[i+1][0]
						grid[i+1][0] = grid[i][j]
					} else if j < len(grid[i]) {
						swap = grid[i][j+1]
						grid[i][j+1] = grid[i][j]
					}
				}
			}
		}
	}

	return grid
}

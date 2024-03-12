package kata

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}

	var sum int
	for i := 0; i < len(grid[0]); i++ {
		var max int
		for j := len(grid) - 1; j >= 0; j-- {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}
		sum += max
	}
	return sum
}

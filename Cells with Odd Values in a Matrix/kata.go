package kata

func oddCells(n int, m int, indices [][]int) int {
	if n == 0 {
		return 0
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	for _, idx := range indices {
		n := idx[0]
		m := idx[1]
		for j := 0; j < len(matrix[n]); j++ {
			matrix[n][j]++
		}
		for i := 0; i < len(matrix); i++ {
			matrix[i][m]++
		}
	}

	var c int
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j]%2 == 1 {
				c++
			}
		}
	}
	return c
}

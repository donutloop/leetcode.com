package kata

func diagonalSum(mat [][]int) int {

	var sum int
	var i, j int
	l := len(mat) - 1
	for i < len(mat) {
		sum += mat[i][j]
		if l == j {
			l--
			i++
			j++
			continue
		}
		sum += mat[i][l]
		l--
		i++
		j++
	}

	return sum
}

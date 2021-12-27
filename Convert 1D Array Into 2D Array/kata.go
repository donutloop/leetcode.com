package kata

func construct2DArray(original []int, m int, n int) [][]int {

	if m*n != len(original) {
		return [][]int{}
	}

	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	k := n
	i := 0
	j := 0
	for _, b := range original {
		matrix[j][i] = b
		k--
		if k == 0 {
			j++
			k = n
			i = 0
		} else {
			i++
		}
	}

	return matrix

}

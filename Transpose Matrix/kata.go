package kata

func transpose(matrix [][]int) [][]int {

	transposeMatrix := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		transposeMatrix[i] = make([]int, len(matrix))
	}

	var k, l int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			transposeMatrix[j][i] = matrix[i][j]
			k++
		}
		k = 0
		l++
	}

	return transposeMatrix
}

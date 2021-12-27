package kata

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		k := len(A[i]) - 1
		for j := 0; j < len(A[i])/2; j++ {
			A[i][j], A[i][k] = A[i][k], A[i][j]
			k--
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] = A[i][j] ^ 1
		}
	}
	return A
}

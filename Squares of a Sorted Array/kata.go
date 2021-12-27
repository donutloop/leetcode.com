package kata

import "sort"

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	return A
}

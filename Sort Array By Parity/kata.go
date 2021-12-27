package kata

import "sort"

func sortArrayByParity(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		return A[i]%2 == 0
	})
	return A
}

package kata

import "sort"

func minIncrementForUnique(A []int) int {

	sort.Ints(A)
	var count int
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			A[i]++
			count++
			continue
		}
		if A[i-1] > A[i] {
			diff := A[i-1] + 1 - A[i]
			A[i] += diff
			count += diff
			continue
		}
	}

	return count
}

package kata

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	var i int
	for K > 0 {
		if A[i] > A[i+1] {
			i++
		}
		if A[i] == 0 {
			break
		}
		A[i] = -1 * A[i]
		K--
	}

	var sum int
	for _, a := range A {
		sum += a
	}

	return sum
}

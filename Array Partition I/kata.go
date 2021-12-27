package kata

import "sort"

func arrayPairSum(a []int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var sum int
	for i := 0; i < len(a); i = i + 2 {
		sum += a[i]
	}
	return sum
}

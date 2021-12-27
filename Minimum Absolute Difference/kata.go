package kata

import "sort"

func minimumAbsDifference(arr []int) [][]int {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	globalMin := -1
	for i := 1; i < len(arr); i++ {
		min := arr[i] - arr[i-1]
		if globalMin == -1 || min < globalMin {
			globalMin = min
		}
	}

	pairs := make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		min := arr[i] - arr[i-1]
		if globalMin == min {
			pairs = append(pairs, []int{arr[i-1], arr[i]})
		}
	}

	return pairs
}

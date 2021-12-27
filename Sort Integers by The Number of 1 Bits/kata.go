package kata

import (
	"sort"
	"strconv"
)

func sortByBits(arr []int) []int {
	mapping := make(map[int]int)
	for _, n := range arr {
		if _, ok := mapping[n]; !ok {
			count := countBits(int64(n))
			mapping[n] = count
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if mapping[arr[i]] == mapping[arr[j]] {
			return arr[i] < arr[j]
		}
		return mapping[arr[i]] < mapping[arr[j]]
	})
	return arr
}

func countBits(n int64) int {
	var count int
	b := strconv.FormatInt(n, 2)
	for _, char := range b {
		if char == '1' {
			count++
		}
	}
	return count
}

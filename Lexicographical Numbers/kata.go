package kata

import (
	"sort"
	"strconv"
)

func lexicalOrder(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i]) < strconv.Itoa(nums[j])
	})
	return nums
}

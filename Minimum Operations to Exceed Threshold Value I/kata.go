package kata

import "sort"

func minOperations(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var countOfsmaller int
	for _, num := range nums {
		if num >= k {
			break
		}
		countOfsmaller++
	}

	return countOfsmaller
}

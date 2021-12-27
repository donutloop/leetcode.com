package kata

import "sort"

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)-k]
}

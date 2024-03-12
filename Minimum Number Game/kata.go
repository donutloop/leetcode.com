package kata

import "sort"

func numberGame(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-1; i = i + 2 {
		if i+1 > len(nums) {
			continue
		}
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}

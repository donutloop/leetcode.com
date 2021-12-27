package kata

import "sort"

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		diff := nums[i-1] - nums[i]
		if diff == 0 {
			return nums[i]
		}
	}
	return -1
}

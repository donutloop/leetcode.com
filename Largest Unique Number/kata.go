package Largest_Unique_Number

import "sort"

func largestUniqueNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 1; i < len(nums); i++ {
		if i > 1 && (nums[i-1] == nums[i-2]) {
			continue
		} else if nums[i-1] == nums[i] {
			i++
			continue
		} else {
			return nums[i-1]
		}
	}

	if len(nums)%2 == 1 {
		return nums[len(nums)-1]
	}

	return -1
}

package Longest_Consecutive_Sequence

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	sort.Ints(nums)

	if len(nums) == 2 {
		if nums[1]-nums[0] == 1 {
			return 2
		}
		return 1
	}

	globalMax := 1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 0 {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			max++
		} else {
			if max > globalMax {
				globalMax = max
			}
			max = 1
		}
	}

	if max > globalMax {
		globalMax = max
	}

	return globalMax
}

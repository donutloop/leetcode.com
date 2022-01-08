package Smallest_Range_I

import "sort"

func smallestRangeI(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}

	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]

	min = min + k

	diff := max - min
	if diff > k {
		diff = k
	}

	return (max - diff) - min
}

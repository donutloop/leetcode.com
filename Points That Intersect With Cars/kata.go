package main

import "sort"

func numberOfPoints(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	var left, right = nums[0][0], nums[0][1]
	var sum = 0
	for _, num := range nums[1:] {
		if num[0] > left && right < num[0] {
			sum += right - (left - 1)
			right = num[1]
			left = num[0]
		} else if num[1] > right {
			right = num[1]
		}
	}

	sum += right - (left - 1)

	return sum
}

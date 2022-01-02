package Largest_Perimeter_Triangle

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i >= 2; i-- {
		a := nums[i-1] + nums[i]
		b := nums[i-2] + nums[i-1]
		c := nums[i-2] + nums[i]
		if a > nums[i-2] && b > nums[i] && c > nums[i-1] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}

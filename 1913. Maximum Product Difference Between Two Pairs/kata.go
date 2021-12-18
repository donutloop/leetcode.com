package _913__Maximum_Product_Difference_Between_Two_Pairs

import "sort"

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1]
}

package Minimize_Product_Sum_of_Two_Arrays

import "sort"

func minProductSum(nums1 []int, nums2 []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums1)))
	sort.Ints(nums2)
	var sum int
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i] * nums2[i]
	}
	return sum
}

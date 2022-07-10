package Find_the_Kth_Largest_Integer_in_the_Array

import "sort"

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) == len(nums[j]) && nums[i] > nums[j] {
			return true
		} else if len(nums[i]) > len(nums[j]) {
			return true
		}
		return false
	})
	return nums[k-1]
}

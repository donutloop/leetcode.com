package Divide_Array_Into_Equal_Pairs

import "sort"

func divideArray(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i = i + 2 {
		if nums[i] != nums[i-1] {
			return false
		}
	}

	return true
}

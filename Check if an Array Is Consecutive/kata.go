package Check_if_an_Array_Is_Consecutive

import "sort"

func isConsecutive(nums []int) bool {

	sort.Ints(nums)

	min, max := nums[0], nums[0]+len(nums)-1

	for i := 1; i < len(nums); i++ {
		if (nums[i] >= min && nums[i] <= max) && (nums[i-1] >= min && nums[i-1] <= max) {
			if nums[i]-nums[i-1] != 1 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

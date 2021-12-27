package _1__Search_in_Rotated_Sorted_Array_II

func search(nums []int, target int) bool {
	if len(nums) == 1 && nums[0] == target {
		return true
	} else if len(nums) == 1 && nums[0] != target {
		return false
	}
	if len(nums) == 0 {
		return false
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return true
	}
	ok := search(nums[:mid], target)
	if ok {
		return true
	}
	ok = search(nums[mid:], target)
	if ok {
		return true
	}
	return false
}

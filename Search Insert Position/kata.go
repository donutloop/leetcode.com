package kata

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target < nums[0] {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if target < nums[i] {
			return i
		}
		if nums[i] == target {
			return i
		}
	}
	return -1
}

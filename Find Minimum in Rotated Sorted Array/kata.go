package kata

func findMin(nums []int) int {
	return search(nums, -1)
}

func search(nums []int, min int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if min == -1 {
		min = nums[len(nums)/2]
	}
	r := search(nums[len(nums)/2:], min)
	if r < min {
		min = r
	}
	l := search(nums[:len(nums)/2], min)
	if l < min {
		min = l
	}
	return min
}

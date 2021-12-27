package _54__Find_Minimum_in_Rotated_Sorted_Array_II

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2
	min := nums[mid]
	leftMin := search(nums[mid:], min)
	if leftMin < min {
		min = leftMin
	}

	rightMin := search(nums[:mid], min)
	if rightMin < min {
		min = rightMin
	}

	return min
}

func search(nums []int, min int) int {
	mid := len(nums) / 2
	if nums[mid] < min {
		min = nums[mid]
	}
	if len(nums) == 1 {
		return min
	}
	leftMin := search(nums[mid:], min)
	if leftMin < min {
		min = leftMin
	}
	rightMin := search(nums[:mid], min)
	if rightMin < min {
		min = rightMin
	}
	return min
}

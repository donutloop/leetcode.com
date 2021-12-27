package kata

func maxAscendingSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var globalMax int
	var localMax int

	localMax = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			localMax += nums[i]
		} else {
			if localMax > globalMax {
				globalMax = localMax
			}
			localMax = nums[i]
		}
	}

	if localMax > globalMax {
		globalMax = localMax
	}

	return globalMax
}

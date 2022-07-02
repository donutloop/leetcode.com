package Min_Max_Game

func minMaxGame(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	for len(nums) != 1 {
		var j int
		for i := 0; i < len(nums)/2; i++ {
			if i%2 == 0 {
				nums[j] = min(nums[2*i], nums[2*i+1])
			} else {
				nums[j] = max(nums[2*i], nums[2*i+1])
			}
			j++
		}
		nums = nums[:len(nums)/2]
	}

	return nums[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

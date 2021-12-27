package kata

func arraySign(nums []int) int {
	var negativ int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			negativ++
		}
	}

	if negativ%2 == 0 {
		return 1
	} else if negativ%2 == 1 {
		return -1
	}

	return 1
}

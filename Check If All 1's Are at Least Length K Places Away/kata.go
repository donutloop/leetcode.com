package kata

func kLengthApart(nums []int, k int) bool {
	var currentK *int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if currentK == nil {
				currentK = new(int)
			} else {
				if *currentK < k {
					return false
				}
				*currentK = 0
			}
			continue
		}
		if currentK != nil {
			*currentK++
		}
	}

	return true
}

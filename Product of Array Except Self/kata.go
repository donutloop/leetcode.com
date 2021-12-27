package kata

func productExceptSelf(nums []int) []int {
	sum := 0
	sumZero := 0
	first := true
	zeros := 0
	for _, n := range nums {
		if n == 0 {
			sumZero *= n
			zeros++
			if zeros > 1 {
				sum *= n
			}
			continue
		}
		if first {
			sum = 1
			sumZero = 1
			first = false
		}

		sum *= n
		sumZero *= n
	}
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = sum
			continue
		}
		nums[i] = sumZero / nums[i]
	}
	return nums
}

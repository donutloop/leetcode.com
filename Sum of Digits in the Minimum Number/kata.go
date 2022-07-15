package kata

func sumOfDigits(nums []int) int {
	min := -1
	for _, n := range nums {
		if min == -1 || min > n {
			min = n
		}
	}

	var sum int
	for {
		r := min % 10
		min = min / 10
		sum += r
		if min == 0 {
			break
		}
	}

	if sum%2 == 0 {
		return 1
	}
	return 0
}

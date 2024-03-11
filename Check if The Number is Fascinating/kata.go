package kata

func isFascinating(n int) bool {
	var stats = [9]int{}
	for i := 1; i <= 3; i++ {
		ok := countDigits(n*i, &stats)
		if !ok {
			return false
		}
	}
	return assert(&stats)
}

func countDigits(n int, stats *[9]int) bool {
	for n > 0 {
		var digit = n % 10
		n = n / 10
		if digit == 0 {
			return false
		}
		digit = digit - 1
		currentCount := stats[digit]
		if currentCount > 0 {
			return false
		}
		stats[digit]++
	}
	return true
}

func assert(stats *[9]int) bool {
	for _, count := range stats {
		if count == 0 || count > 1 {
			return false
		}
	}
	return true
}

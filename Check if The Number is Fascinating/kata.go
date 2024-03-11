package kata

func isFascinating(n int) bool {
	var stats = [9]int{}
	var tracker int = 0
	for i := 1; i <= 3; i++ {
		ok := countDigits(n*i, &stats, &tracker)
		if !ok {
			return false
		}
	}
	return true
}

func countDigits(n int, stats *[9]int, tracker *int) bool {
	for n > 0 {
		var digit = n % 10
		n = n / 10
		digit = digit - 1
		if digit == -1 || stats[digit] > 0 {
			return false
		}
		stats[digit]++
		*tracker++
	}
	return *tracker%3 == 0
}

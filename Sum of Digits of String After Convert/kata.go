package kata

func getLucky(s string, k int) int {

	var sum uint64
	for i := len(s) - 1; i >= 0; i-- {
		k := uint64(s[i] - 96)

		if k > 9 {
			for k != 0 {
				r := k % 10
				k /= 10
				sum += r
			}

		} else {
			sum += k
		}
	}

	for k-1 > 0 {
		newSum := sum
		sum = 0
		for newSum != 0 {
			r := newSum % 10
			newSum /= 10
			sum += r
		}
		k--
	}
	return int(sum)
}

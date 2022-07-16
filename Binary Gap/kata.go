package kata

func binaryGap(n int) int {
	var globalMax int
	var max int
	for n != 0 {
		r := n % 2
		n = n / 2
		if r == 0 && max == 0 {
			continue
		} else if r == 1 && max == 0 {
			max++
		} else if r == 1 && max != 0 {
			if max > globalMax {
				globalMax = max
			}
			max = 1
		} else {
			max++
		}
	}

	return globalMax
}

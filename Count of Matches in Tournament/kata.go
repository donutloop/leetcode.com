package kata

func numberOfMatches(n int) int {

	var totalMatches int
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
			totalMatches += n
		} else {
			n = (n - 1) / 2
			totalMatches += n
			n = n + 1
		}
	}

	return totalMatches
}

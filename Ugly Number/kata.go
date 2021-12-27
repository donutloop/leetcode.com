package kata

func isUgly(num int) bool {
	if num == 1 {
		return true
	}

	if num <= 0 {
		return false
	}

	for {
		if num == 3 || num == 5 || num == 2 {
			return true
		}

		if num%2 == 0 {
			num = num / 2
			continue
		} else if num%3 == 0 {
			num = num / 3
			continue
		} else if num%5 == 0 {
			num = num / 5
			continue
		}

		break
	}

	return false
}

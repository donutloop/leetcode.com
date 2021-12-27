package kata

func isHappy(n int) bool {
	var ok bool
	steps := make(map[int]int)
	for n > 0 {
		n = sumFunc(n)
		if n == 1 {
			ok = true
			break
		}
		_, found := steps[n]
		if found {
			ok = false
			break
		}
		steps[n] = 0
	}

	return ok
}

func sumFunc(n int) int {
	var sum int
	for n > 0 {
		r := n % 10
		n = n / 10
		sum = sum + (r * r)
	}
	return sum
}

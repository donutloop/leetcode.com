package kata

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	var digitCount int
	for {
		num, digitCount = sumFunc(num)
		if digitCount == 1 {
			return num
		}
	}
}

func sumFunc(num int) (int, int) {
	var s int
	var c int
	for num > 0 {
		n := num % 10
		num = num / 10
		s += n
		c++

	}
	return s, c
}

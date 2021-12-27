package kata

// 0, 0, 1, 1, 2, 4, 7, 13, 24, 44, 81, 149, 274, 504, 927, 1705, 3136, 5768, 10609, 19513, 35890, 66012

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	solver := eq()
	var k int
	for i := 0; i < n-2; i++ {
		k = solver()
	}
	return k
}

func eq() func() int {
	x1, x2, x3 := 0, 1, 1
	return func() int {
		x1, x2, x3 = x2, x3, x1+x2+x3
		return x3
	}
}

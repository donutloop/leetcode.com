package kata

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	i := 1
	c := 0
	for n > 0 {
		n = n - i
		if n < 0 {
			break
		}
		c++
		i++
	}
	return c
}

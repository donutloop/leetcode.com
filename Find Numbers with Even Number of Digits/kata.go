package kata

func findNumbers(nums []int) int {
	var c int
	for _, number := range nums {
		i := 0
		for number != 0 {
			i++
			number /= 10
		}
		if i%2 == 0 {
			c++
		}
	}

	return c
}

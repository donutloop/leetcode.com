package kata

func hammingDistance(x int, y int) int {
	var count int
	for y != 0 || x != 0 {

		b2 := 0
		if x != 0 {
			b2 = x % 2
			x = x / 2
		}

		b1 := 0
		if y != 0 {
			b1 = y % 2
			y = y / 2
		}

		if b1 != b2 {
			count++
		}
	}
	return count
}

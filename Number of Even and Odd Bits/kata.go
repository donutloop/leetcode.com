package kata

func evenOddBit(n int) []int {
	var i int
	var odd int
	var even int
	for n > 0 {
		b := n % 2
		n = n / 2
		if b == 1 {
			if i%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		i++
	}
	return []int{even, odd}
}

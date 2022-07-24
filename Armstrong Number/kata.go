package kata

import "math"

func isArmstrong(n int) bool {

	l := n
	var c int
	for l != 0 {
		l /= 10
		c++
	}

	var m int
	k := n
	for n != 0 {

		r := n % 10
		n /= 10

		m += int(math.Pow(float64(r), float64(c)))
		if m > k {
			return false
		}
	}

	return m == k
}

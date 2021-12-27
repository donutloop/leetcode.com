package kata

import "math/bits"

func sumBase(n int, k int) int {

	var sum int
	for n != 0 {
		q, r := bits.Div(0, uint(n), uint(k))
		n = int(q)
		sum += int(r)
	}

	return sum
}

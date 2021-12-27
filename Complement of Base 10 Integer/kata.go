package kata

import "math"

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	} else if N == 1 {
		return 0
	}

	var sum int
	var k int
	for N > 0 {
		b := N % 2
		N = N / 2
		if b == 0 {
			b = 1
		} else {
			b = 0
		}

		sum += b * int(math.Pow(2, float64(k)))
		k++
	}

	return sum
}

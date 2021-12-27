package kata

import "math"

func isPowerOfTwo(n int) bool {
	if 0 == n {
		return false
	}
	k := math.Log2(float64(n))
	return math.Floor(k) == k
}

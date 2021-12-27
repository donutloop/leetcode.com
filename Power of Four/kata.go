package kata

import "math"

func isPowerOfFour(num int) bool {
	if 0 == num {
		return false
	}
	k := math.Log(float64(num)) / math.Log(4)
	return math.Floor(k) == k
}

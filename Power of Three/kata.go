package kata

import "math"

func isPowerOfThree(num int) bool {
	if 0 == num {
		return false
	}
	k := math.Log(float64(num)) / math.Log(3)
	n1 := math.Pow(3, float64(int(k)))

	return int(n1) == num
}

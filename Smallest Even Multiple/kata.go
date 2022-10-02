package Smallest_Even_Multiple

import "math"

func smallestEvenMultiple(n int) int {
	k := float64(n) / 2
	u := math.Floor(k)
	if u == k {
		return n
	}
	return 2 * n
}

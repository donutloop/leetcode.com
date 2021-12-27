package kata

import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

	var c int
	for _, n1 := range arr1 {
		var has bool
		for _, n2 := range arr2 {
			newDistance := math.Abs(float64(n1) - float64(n2))
			if newDistance <= float64(d) {
				has = true
				break
			}
		}

		if !has {
			c++
		}
	}

	return c
}

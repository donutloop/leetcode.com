package Find_Greatest_Common_Divisor_of_Array

import "math"

func findGCD(nums []int) int {

	min := int(math.Inf(-1))
	max := int(math.Inf(-1))
	for _, n := range nums {
		if max == int(math.Inf(-1)) || n > max {
			max = n
		}
		if min == int(math.Inf(-1)) || n < min {
			min = n
		}
	}

	for gcd := max; gcd > 0; gcd-- {
		if min%gcd == 0 && max%gcd == 0 {
			return gcd
		}
	}

	return -1
}

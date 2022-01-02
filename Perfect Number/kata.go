package Perfect_Number

import "math"

func checkPerfectNumber(num int) bool {

	var sum float64
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			if num/i == i {
				sum += float64(i)
			} else {
				sum += float64(i) + float64(num)/float64(i)
			}
		}
	}

	return int(sum-float64(num)) == num
}

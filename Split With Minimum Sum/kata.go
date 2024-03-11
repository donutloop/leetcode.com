package kata

import (
	"math"
	"sort"
)

func splitNum(num int) int {
	if num <= 9 {
		return num
	}

	digits := make([]int, int(math.Log10(float64(num)))+1)
	var i int
	for num > 0 {
		var digit = num % 10
		num = num / 10
		digits[i] = digit
		i++
	}

	if len(digits) == 2 {
		return digits[0] + digits[1]
	}

	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})

	a := digits[0]
	b := digits[1]

	var multiplierOfA, multiplierOfB = 10, 10
	for i := 2; i < len(digits); i++ {
		if i%2 == 0 {
			b += digits[i] * multiplierOfB
			multiplierOfB = multiplierOfB * 10
		} else {
			a += digits[i] * multiplierOfA
			multiplierOfA = multiplierOfA * 10
		}
	}

	return a + b
}

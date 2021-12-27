package kata

import "math"

const (
	Negative int = -1
	Positive int = 1
)

func reverse(x int) (sum int) {

	var sign int
	if x < 0 {
		sign = Negative
		x = x * Negative
	} else if x > 0 {
		sign = Positive
	}

	step := 1
	for i := int(math.Log10(float64(x)+1)) - 1; i >= 0; i-- {
		step = step * 10
	}

	for x != 0 {
		n := x % 10
		x = x / 10
		sum += n * step
		step = step / 10
	}

	if sum > math.MaxInt32 {
		return 0
	}

	if sign == Negative {
		sum = sum * sign
	}
	return sum
}

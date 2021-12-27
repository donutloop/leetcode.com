package kata

import "math"

func findComplement(num int) int {
	bin := make([]int, 0)
	for num > 0 {
		k := num % 2
		num = num / 2
		bin = append(bin, k)
	}

	for i := range bin {
		if bin[i] == 1 {
			bin[i] = 0
		} else {
			bin[i] = 1
		}
	}

	num = 0
	for i := 0; i < len(bin); i++ {
		num += bin[i] * int(math.Pow(float64(2), float64(i)))
	}

	return num
}

package kata

import "math"

func sumOfEncryptedInt(nums []int) int {
	var sum int
	for _, num := range nums {
		var digitCount float64
		var max int
		for num > 0 {
			digit := num % 10
			if max < digit {
				max = digit
			}
			num = num / 10
			digitCount++
		}

		sum += max * (1 - int(math.Pow(10, digitCount))) / (1 - 10)
	}

	return sum
}

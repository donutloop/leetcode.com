package kata

func sumOfTheDigitsOfHarshadNumber(x int) int {
	num := x
	var sum int
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	if x%sum == 0 {
		return sum
	}
	return -1
}

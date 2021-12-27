package kata

func subtractProductAndSum(number int) int {

	multiplySum := 1
	plusSum := 0
	for number > 0 {
		val := number % 10
		multiplySum *= val
		plusSum += val
		number = number / 10
	}

	return multiplySum - plusSum
}

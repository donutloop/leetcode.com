package Count_Operations_to_Obtain_Zero

func countOperations(num1 int, num2 int) int {

	if num1 == 0 || num2 == 0 {
		return 0
	}

	var c int
	for {
		a := num1
		num1 = max(num1, num2)
		num2 = min(a, num2)

		num1 = num1 - num2

		c++
		if num1 == 0 {
			break
		}
	}

	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

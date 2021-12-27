package kata

func convertToBase7(num int) string {

	if num == 0 {
		return "0"
	}

	var neg bool
	if num < 0 {
		num = num * -1
		neg = true
	}

	convertedNum := make([]byte, 0)
	for num > 0 {
		rem := num % 7
		num = num / 7
		rem = rem + 48
		convertedNum = append(convertedNum, byte(rem))
	}

	for i, j := 0, len(convertedNum)-1; i < j; i, j = i+1, j-1 {
		convertedNum[i], convertedNum[j] = convertedNum[j], convertedNum[i]
	}

	if neg {
		return "-" + string(convertedNum)
	}

	return string(convertedNum)
}

package kata

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1 = padding(num2, num1)
		return sumFunc(num1, num2)
	} else if len(num1) > len(num2) {
		num2 = padding(num1, num2)
		return sumFunc(num1, num2)
	}
	return sumFunc(num1, num2)
}

func sumFunc(num1, num2 string) string {
	num3 := []byte(num1)
	var lastReminder byte
	for i := len(num2) - 1; i >= 0; i-- {
		sum := ((num1[i] - 48) + (num2[i] - 48)) + lastReminder
		if sum > 9 {
			diff := sum - byte(9)
			if diff > 1 {
				num3[i] = diff - 1 + 48
			} else {
				num3[i] = 48
			}
			lastReminder = byte(1)
		} else {
			num3[i] = sum + 48
			lastReminder = byte(0)
		}
	}
	if lastReminder > 0 {
		return "1" + string(num3)
	}

	return string(num3)
}

func padding(num2, num1 string) string {
	diff := len(num2) - len(num1)
	for diff > 0 {
		num1 = "0" + num1
		diff--
	}
	return num1
}

package Check_if_Number_Has_Equal_Digit_Count_and_Digit_Value

func digitCount(num string) bool {

	c := map[rune]int{}
	for _, digit := range num {
		c[digit]++
	}

	for i, digit := range num {
		if c[rune(i+48)] != int(digit-48) {
			return false
		}
	}

	return true
}

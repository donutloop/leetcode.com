package _6__Plus_One

var one []int = []int{1}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return one
	}

	var reminder bool
	if digits[len(digits)-1] == 9 {
		for i := len(digits) - 1; i >= 0; i-- {
			if reminder && (digits[i]+1) == 10 {
				digits[i] = 0
				reminder = true
			} else if reminder {
				digits[i] += 1
				reminder = false
				break
			} else if (digits[i] + 1) == 10 {
				digits[i] = 0
				reminder = true
			}
		}

		if reminder {
			return append(one, digits...)
		}

	} else {
		digits[len(digits)-1] = digits[len(digits)-1] + 1
	}
	return digits
}

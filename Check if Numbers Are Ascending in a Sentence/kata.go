package kata

import "math"

func areNumbersAscending(s string) bool {

	currentNumber := int(math.Inf(-1))
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 48 && s[i] <= 57 {

			n := 1
			newNumber := (int(s[i]) - 48) * n
			i = i - 1
			for ; i >= 0; i-- {

				if s[i] >= 48 && s[i] <= 57 {
					n = n * 10
					newNumber += (int(s[i]) - 48) * n
				} else {
					break
				}

			}

			if currentNumber == int(math.Inf(-1)) {
				currentNumber = newNumber
			} else {
				if currentNumber <= newNumber {
					return false
				}
				currentNumber = newNumber
			}
		}
	}

	return true
}

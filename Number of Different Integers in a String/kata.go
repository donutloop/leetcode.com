package kata

func numDifferentIntegers(word string) int {
	var found int
	count := make(map[string]int)
	digit := []rune{}
	for _, char := range word {
		if char >= 48 && char <= 57 {
			found++
			digit = append(digit, char)
		}
		if char >= 97 && char <= 122 {
			if found >= 1 {
				found = 0

				var i int
				for _, n := range digit {
					if n == '0' {
						i++
					} else {
						break
					}
				}
				count[string(digit[i:])] = 0
				digit = []rune{}
			}
		}
	}

	if found >= 1 {
		var i int
		for _, n := range digit {
			if n == '0' {
				i++
			} else {
				break
			}
		}
		count[string(digit[i:])] = 0
	}

	return len(count)
}

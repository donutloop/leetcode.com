package _25__Valid_Palindrome

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {

		ok, addI := isValidChar(s[i])
		if !ok {
			continue
		}

		ok, addJ := isValidChar(s[j])
		if !ok {
			j--
			i--
			continue
		}

		if (int(s[i]) + addI) != (int(s[j]) + addJ) {
			return false
		}
		j--
	}
	return true
}

func isValidChar(char byte) (bool, int) {
	if char >= 97 && char <= 122 {
		return true, 0
	} else if char >= 65 && char <= 90 {
		return true, 32
	} else if char >= 48 && char <= 57 {
		return true, 0
	}
	return false, 0
}

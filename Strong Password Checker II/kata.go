package Strong_Password_Checker_II

var specialCharacters = map[byte]bool{'!': true, '@': true, '#': true, '$': true, '%': true, '^': true, '&': true, '*': true, '(': true, ')': true, '-': true, '+': true}

func strongPasswordCheckerII(password string) bool {

	if len(password) < 8 {
		return false
	}

	var lowerCase bool
	var upperCase bool
	var digit bool
	var specialCharacter bool
	var adjacent bool

	j := 1
	for i := 0; i < len(password); i++ {

		if j < len(password) && password[j] == password[i] {
			adjacent = true
		}

		if specialCharacters[password[i]] {
			specialCharacter = true
		} else if password[i] >= 48 && password[i] <= 57 {
			digit = true
		} else if password[i] >= 97 && password[i] <= 122 {
			lowerCase = true
		} else if password[i] >= 65 && password[i] <= 90 {
			upperCase = true
		}

		j++

	}

	return lowerCase && upperCase && digit && specialCharacter && !adjacent
}

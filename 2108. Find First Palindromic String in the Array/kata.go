package _108__Find_First_Palindromic_String_in_the_Array

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[j] != s[i] {
			return false
		}
		j--
	}
	return true
}

package _8__Length_of_Last_Word

func lengthOfLastWord(s string) (lastWordCounter int) {
	if s == "" {
		return 0
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && lastWordCounter != 0 {
			return
		}
		if s[i] != ' ' {
			lastWordCounter++
		}
	}
	return
}

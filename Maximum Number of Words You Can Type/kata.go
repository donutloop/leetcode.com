package Maximum_Number_of_Words_You_Can_Type

func canBeTypedWords(text string, brokenLetters string) int {

	chars := make(map[rune]bool)
	for _, char := range brokenLetters {
		chars[char] = true
	}

	var goodWordsCounter int

	nextWord := false
	for i, char := range text {
		if nextWord {
			if char == 32 {
				nextWord = false
			}
			continue
		}
		if chars[char] {
			nextWord = true
			continue
		}

		if char == 32 || (len(text)-1) == i {
			goodWordsCounter++
		}
	}

	return goodWordsCounter
}

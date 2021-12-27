package kata

func commonChars(words []string) []string {

	chars := []string{}

	counter := make([][26]int, len(words))

	for w, word := range words {
		for i := 0; i < len(word); i++ {
			counter[w][word[i]-97]++
		}
	}

outerLoop:
	for j := 0; j < 26; j++ {
		min := int(-1)
		for i := 0; i < len(words); i++ {
			counter := int(counter[i][j])
			if counter == 0 {
				continue outerLoop
			}

			if min == -1 || min > counter {
				min = counter
			}
		}

		for i := 0; i < min; i++ {
			chars = append(chars, string([]byte{byte(j + 97)}))
		}

	}

	return chars
}

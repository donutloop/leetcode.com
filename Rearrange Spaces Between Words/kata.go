package kata

func reorderSpaces(text string) string {

	var words [][]byte
	var spaceCount int
	var word []byte
	for i := 0; i < len(text); i++ {

		if text[i] == 32 || len(text)-1 == i {

			if len(text)-1 == i && text[i] != 32 {
				word = append(word, text[i])
			} else {
				spaceCount++
			}

			if len(word) != 0 {
				words = append(words, word)
				word = []byte{}
			}

		} else {
			word = append(word, text[i])
		}
	}

	if len(words) == 1 && spaceCount == 0 {
		return text
	}

	var inbetweenCount int
	if len(words) == 1 {
		inbetweenCount = 1
	} else {
		inbetweenCount = (len(words) - 1)
	}

	mod := spaceCount % inbetweenCount
	spaceCount = spaceCount / inbetweenCount

	var space string
	for i := 0; i < spaceCount; i++ {
		space += " "
	}

	var restSpace string
	for i := 0; i < mod; i++ {
		restSpace += " "
	}

	if len(words) == 1 {
		return string(words[0]) + space
	}

	text = ""
	for i, word := range words {
		if len(words)-1 == i {
			text += string(word) + restSpace
		} else {

			text += string(word) + space
		}

	}

	return text

}

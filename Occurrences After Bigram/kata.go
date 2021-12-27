package kata

func findOcurrences(text string, first string, second string) []string {

	var j int
	matchedFirst := -1
	matchedSecond := -1
	wordCount := 0
	words := make([]string, 0)
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' || len(text)-1 == i {
			if len(text)-1 == i {
				i = i + 1
			}
			wordCount++

			if matchedFirst == -1 && text[j:i] == first {
				matchedFirst = wordCount
				j = i + 1
				continue
			} else if wordCount-1 == matchedFirst && text[j:i] == second {
				matchedSecond = wordCount
				j = i + 1
				continue
			} else if wordCount-1 == matchedSecond {
				words = append(words, text[j:i])
				matchedSecond = -1
				matchedFirst = -1
				if text[j:i] == first {
					matchedFirst = wordCount
				}
				j = i + 1
				continue
			}

			matchedFirst = -1
			if text[j:i] == first {
				matchedFirst = wordCount
			}
			matchedSecond = -1
			j = i + 1
		}
	}

	return words
}

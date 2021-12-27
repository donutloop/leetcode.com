package kata

func findAndReplacePattern(words []string, pattern string) []string {
	cleanWords := cleanWords(words)
	pattern = cleanWord(pattern)

	var matched []string
	for i, w := range cleanWords {
		if w == pattern {
			matched = append(matched, words[i])
		}
	}

	return matched
}

func cleanWords(words []string) []string {
	cws := make([]string, len(words))
	for i, w := range words {
		cws[i] = cleanWord(w)
	}
	return cws
}

func cleanWord(w string) string {
	cleanWord := "A"
	char := make(map[byte]byte)
	lastChar := w[0]
	char[w[0]] = 65
	for i := 1; i < len(w); i++ {
		if lastChar == w[i] {
			cleanWord += string(cleanWord[i-1])
		} else if e, ok := char[w[i]]; ok {
			cleanWord += string(e)
		} else {
			cleanWord += string(cleanWord[i-1] + 1)
		}
		lastChar = w[i]
		char[w[i]] = cleanWord[i-1] + 1
	}
	return cleanWord
}

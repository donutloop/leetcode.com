package kata

func makeEqual(words []string) bool {
	charStats := [26]int{}
	countOfchars := 0
	for _, word := range words {
		for _, codepoint := range word {
			var normalisedCodepoint = codepoint - 97
			charStats[normalisedCodepoint]++
			countOfchars++
		}
	}

	if countOfchars%len(words) == 1 {
		return false
	}

	for _, oneCharStat := range charStats {
		if oneCharStat > 0 && oneCharStat%len(words) != 0 {
			return false
		}
	}

	return true
}

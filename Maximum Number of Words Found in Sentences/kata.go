package kata

func mostWordsFound(sentences []string) int {
	var globalMax int
	for _, sentence := range sentences {
		var max int
		for i := 0; i < len(sentence); i++ {
			if sentence[i] == ' ' {
				max++
			}
		}
		max++
		if max > globalMax {
			globalMax = max
		}
	}
	return globalMax
}

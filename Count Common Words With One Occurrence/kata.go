package kata

func countWords(words1 []string, words2 []string) int {

	type wordCounter struct {
		A, B int
	}

	counter := make(map[string]wordCounter)

	for _, word := range words1 {
		wordCounter := counter[word]
		wordCounter.A++
		counter[word] = wordCounter
	}

	for _, word := range words2 {
		wordCounter := counter[word]
		wordCounter.B++
		counter[word] = wordCounter
	}

	var total int
	for _, wordCounter := range counter {
		if wordCounter.A == 1 && wordCounter.B == 1 {
			total++
		}

	}

	return total
}

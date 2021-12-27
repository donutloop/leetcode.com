package kata

func isPrefixOfWord(sentence string, searchWord string) int {
	var wordCounter int
	var matched bool
	var j int
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' || i == len(sentence)-1 {
			if i == len(sentence)-1 {
				i = i + 1
			}
			wordCounter++
			if len(sentence[j:i]) >= len(searchWord) {
				if sentence[j:j+len(searchWord)] == searchWord {
					matched = true
					break
				}
			}
			j = i + 1
		}

	}

	if !matched {
		return -1
	}

	return wordCounter
}

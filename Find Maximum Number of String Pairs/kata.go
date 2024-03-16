package kata

func maximumNumberOfStringPairs(words []string) int {
	pair := make(map[string]int, 0)
	var c int
	for i := 0; i < len(words); i++ {
		pair[words[i]]++
		reversedWord := string(words[i][1]) + string(words[i][0])
		if reversedWord == words[i] {
			continue
		}
		pair[reversedWord]++
		if pair[reversedWord] == 2 {
			pair[reversedWord] = pair[reversedWord] - 2
			c++
		}
	}

	return c
}

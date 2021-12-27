package kata

func mergeAlternately(word1 string, word2 string) string {

	var max int
	var word2IsLongest bool
	if len(word1) > len(word2) {
		max = len(word2)
	} else {
		max = len(word1)
		word2IsLongest = true
	}

	word3 := make([]byte, len(word2)+len(word1))
	i := 0
	j := 0
	for ; i < max; i++ {
		word3[j] = word1[i]
		j++
		word3[j] = word2[i]
		j++
	}

	if max == len(word1) && max == len(word2) {
		return string(word3)
	}

	if word2IsLongest {
		for ; i < len(word2); i++ {
			word3[j] = word2[i]
			j++
		}
	} else {
		for ; i < len(word1); i++ {
			word3[j] = word1[i]
			j++
		}
	}

	return string(word3)
}

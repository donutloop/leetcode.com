package kata

func arrayStringsAreEqual(word1 []string, word2 []string) bool {

	var i int
	var j int
	var a int
	var b int
	for a < len(word1) && b < len(word2) {

		if word1[a][i] != word2[b][j] {
			return false
		}

		if a < len(word1) && len(word1[a])-1 == i {
			i = 0
			a++
			if b < len(word2) && len(word2[b])-1 == j {
				j = 0
				b++
				continue
			}

			j++
			continue
		}

		if b < len(word2) && len(word2[b])-1 == j {
			j = 0
			b++

			if a < len(word1) && len(word1[a])-1 == i {
				i = 0
				a++
				continue
			}

			i++
			continue
		}

		i++
		j++
	}

	if i != 0 || j != 0 {
		return false
	}

	return true
}

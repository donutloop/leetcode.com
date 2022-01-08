package Word_Pattern

func wordPattern(pattern string, s string) bool {

	var k int
	p := pattern[k]

	patternMapping := make(map[string]string)

	var j int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || (len(s)-1) == i {
			if (len(s) - 1) == i {
				i = i + 1
			}
			word, ok := patternMapping[string(p-64)]
			if ok {
				if s[j:i] != word {
					return false
				}
			} else {
				_, ok := patternMapping[s[j:i]]
				if ok {
					return false
				}

				patternMapping[string(p-64)] = s[j:i]
				patternMapping[s[j:i]] = string(p - 64)
			}

			j = i + 1

			if (k + 1) < len(pattern) {
				k++
				p = pattern[k]
			} else {
				k++
			}
		}
	}

	return k == len(pattern)
}

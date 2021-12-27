package kata

func reverseVowels(s string) string {
	b := []byte(s)
	j := len(b) - 1
	i := 0
	for j > i {
		if isVowel(b[i]) {
			if isVowel(b[j]) {
				b[i], b[j] = b[j], b[i]
				i++
			}
			j--
			continue
		}
		i++
	}
	return string(b)
}

func isVowel(char byte) bool {
	switch char {
	case 'a', 'i', 'e', 'o', 'u', 'A', 'I', 'E', 'O', 'U':
		return true
	}
	return false
}

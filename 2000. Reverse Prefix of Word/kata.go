package _000__Reverse_Prefix_of_Word

func reversePrefix(word string, ch byte) string {

	j := -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			j = i
			break
		}
	}

	if j == -1 {
		return word
	}

	b := []byte(word)
	k := 0
	for i := j; i > (j / 2); i-- {
		b[i], b[k] = b[k], b[i]
		k++
	}

	return string(b)
}

package kata

const whitespace = ' '

func truncateSentence(s string, k int) string {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == whitespace {
			count++
		}

		if count == k {
			return s[:i]
		}
	}
	return s
}

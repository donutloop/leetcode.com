package kata

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return sum(firstWord)+sum(secondWord) == sum(targetWord)
}

func sum(s string) int {
	var n int
	m := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 'a' {
			n += (int(s[i]) - 97) * m
		}
		m = m * 10
	}
	return n
}

package kata

func restoreString(s string, indices []int) string {
	newS := make([]byte, len(s))
	for i, j := range indices {
		newS[j] = s[i]
	}
	return string(newS)
}

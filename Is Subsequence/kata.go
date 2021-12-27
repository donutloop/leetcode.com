package kata

func isSubsequence(s string, t string) bool {
	sub := make([]byte, len(s))
	var i int
	for j := 0; j < len(t); j++ {
		if i == len(s) {
			break
		}
		if s[i] == t[j] {
			sub[i] = s[i]
			i++
		}
	}
	if string(sub) == s {
		return true
	}
	return false
}

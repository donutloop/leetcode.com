package kata

func diStringMatch(S string) []int {
	n := make([]int, len(S)+1)
	j := len(S)
	k := 0
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			n[i] = k
			k++
		} else {
			n[i] = j
			j--
		}
	}

	if S[len(S)-1] == 'I' {
		n[len(n)-1] = j
	} else {
		n[len(n)-1] = k
	}
	return n
}

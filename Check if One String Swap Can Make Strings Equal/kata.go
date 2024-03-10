package kata

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s1) {
		return false
	}
	if s1 == s2 {
		return true
	}

	var distance int
	var pairs [][]int
	for i, codepoint := range s1 {
		if codepoint != rune(s2[i]) {
			distance++
			if distance > 2 {
				return false
			}
			pair := []int{int(codepoint), int(s2[i])}
			pairs = append(pairs, pair)
		}
	}

	if len(pairs) == 2 && pairs[0][0] == pairs[1][1] && pairs[0][1] == pairs[1][0] {
		return true
	}

	return false
}

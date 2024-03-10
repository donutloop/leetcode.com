package kata

func minSteps(s string, t string) int {

	type pair struct {
		a, b int
	}

	var tracker = make(map[rune]pair)
	for _, codepoint := range s {
		pair := tracker[codepoint]
		pair.a++
		tracker[codepoint] = pair
	}
	for _, codepoint := range t {
		pair := tracker[codepoint]
		pair.b++
		tracker[codepoint] = pair
	}

	var diffTotal int
	for _, count := range tracker {
		if count.b > count.a {
			diff := abs(count.b - count.a)
			diffTotal += diff
		}
	}

	return diffTotal + len(s) + diffTotal - len(t)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package kata

func checkIfPangram(sentence string) bool {

	var abc [26]int
	var count int
	for _, char := range sentence {
		abc[char-97] += 1
		if abc[char-97] == 1 {
			count++
		}
	}

	if count != 26 {
		return false
	}

	return true
}

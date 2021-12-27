package kata

func detectCapitalUse(word string) bool {
	var capitals int
	var lowers int
	for i := 0; i < len(word); i++ {
		if word[i] >= 65 && word[i] <= 90 {
			capitals++
			if capitals == 2 && lowers >= 1 {
				break
			}
		} else {
			lowers++
		}
	}

	if capitals == 0 || capitals == len(word) || (capitals == 1 && (word[0] >= 65 && word[0] <= 90)) {
		return true
	}
	return false
}

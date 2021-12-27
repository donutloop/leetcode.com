package kata

func countConsistentStrings(allowed string, words []string) int {

	set := make(map[rune]bool)
	for _, char := range allowed {
		set[char] = true
	}

	var count int
	for _, word := range words {
		ok := true
		for _, char := range word {
			if !set[char] {
				ok = false
				break
			}
		}

		if ok {
			count++
		}
	}

	return count
}

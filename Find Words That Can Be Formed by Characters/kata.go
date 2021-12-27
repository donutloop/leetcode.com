package kata

func countCharacters(words []string, chars string) int {

	stats := make(map[rune]int, len(chars))
	for _, c := range chars {
		stats[c]++
	}

	var sum int
outerLoop:
	for _, word := range words {
		if len(word) > len(chars) {
			continue
		}
		currentStats := make(map[rune]int)
		for _, char := range word {
			allowedCount, ok := stats[char]
			if !ok || currentStats[char] >= allowedCount {
				continue outerLoop
			}
			currentStats[char]++
		}
		sum += len(word)
	}

	return sum
}

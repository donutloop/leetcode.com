package Count_Prefixes_of_a_Given_String

func countPrefixes(words []string, s string) int {
	var count int
	for _, w := range words {
		if len(w) > len(s) {
			continue
		}
		if w == s[:len(w)] {
			count++
		}
	}
	return count
}

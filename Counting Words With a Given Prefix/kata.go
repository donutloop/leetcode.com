package Counting_Words_With_a_Given_Prefix

func prefixCount(words []string, pref string) int {
	var c int
	for _, w := range words {
		if len(w) >= len(pref) && w[:len(pref)] == pref {
			c++
		}
	}
	return c
}

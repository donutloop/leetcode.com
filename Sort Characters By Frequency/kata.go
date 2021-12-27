package kata

import "sort"

func frequencySort(s string) string {
	set := make(map[rune]int)
	for _, c := range s {
		set[c]++
	}
	b := []rune(s)
	sort.Slice(b, func(i, j int) bool {
		if set[b[i]] == set[b[j]] && b[i] < b[j] {
			return true
		} else if set[b[i]] > set[b[j]] {
			return true
		}
		return false
	})
	return string(b)
}

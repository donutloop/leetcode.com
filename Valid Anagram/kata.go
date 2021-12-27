package kata

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sb := []byte(s)
	tb := []byte(t)
	sort.Slice(tb, func(i, j int) bool {
		return tb[i] < tb[j]
	})
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	if string(tb) == string(sb) {
		return true
	}
	return false
}

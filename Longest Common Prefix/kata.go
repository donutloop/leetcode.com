package kata

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	min := -1
	minString := ""
	for _, s := range strs {
		if min == -1 || min > len(s) {
			min = len(s)
			minString = s
		}
	}
	for n := len(minString); n >= 0; n-- {
		if Has(strs, minString, n) {
			return minString[:n]
		}
	}
	return ""
}

func Has(strs []string, minString string, n int) bool {
	for _, s := range strs {
		if minString[:n] != s[:n] {
			return false
		}
	}
	return true
}

package kata

func strStr(haystack string, needle string) (idx int) {
	if haystack == needle {
		return 0
	}
	idx = -1
	for i := 0; i < len(haystack); i++ {

		if (i + len(needle)) > len(haystack) {
			return
		}

		if haystack[i:i+len(needle)] == needle {
			idx = i
			return
		}
	}
	return
}

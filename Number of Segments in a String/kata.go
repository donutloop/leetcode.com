package kata

func countSegments(s string) int {
	var c int
	var word int
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word++
		}
		if (s[i] == ' ' || len(s)-1 == i) && word > 0 {
			c++
			word = 0
		}
	}
	return c
}

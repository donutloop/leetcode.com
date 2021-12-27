package kata

func validPalindrome(s string) bool {
	j := len(s) - 1
	Jidx := -1
	Iidx := -1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			Jidx = j
			Iidx = i
			break
		}
		j--
	}

	b := makeValid(s, Iidx)
	validI := IsValid(b)

	b = makeValid(s, Jidx)
	validJ := IsValid(b)

	return validJ || validI
}

func makeValid(s string, idx int) []byte {
	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if i == idx && idx != -1 {
			continue
		}
		b = append(b, s[i])
	}
	return b
}

func IsValid(b []byte) bool {
	valid := true
	j := len(b) - 1
	for i := 0; i < len(b)/2; i++ {
		if b[i] != b[j] {
			valid = false
		}
		j--
	}
	return valid
}

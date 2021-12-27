package kata

func halvesAreAlike(s string) bool {
	aCount := count(s[len(s)/2:])
	bCount := count(s[:len(s)/2])

	if aCount != bCount {
		return false
	}

	return true
}

func count(a string) int {
	aCount := 0
	for i := 0; i < len(a); i++ {
		switch a[i] {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			aCount += 1
		}
	}
	return aCount
}

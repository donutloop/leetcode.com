package kata

const A int = 1
const L int = 2

func checkRecord(s string) bool {
	aCounter := 0
	lCounter := 0
	for i, c := range s {
		if c == 'A' {
			aCounter++
			if aCounter > A {
				return false
			}
		} else if c == 'L' {
			j := i - 1
			if j >= 0 && rune(s[j]) == c {
				lCounter++
				if lCounter > L {
					return false
				}
			} else {
				lCounter = 1
			}
		}
	}
	return true
}

package kata

func balancedStringSplit(s string) int {

	max := 0

	current := s[0]
	seekCounter := 1
	var seek byte
	if current == 'R' {
		seek = 'L'
	} else {
		seek = 'R'
	}

	for i := 1; i < len(s); i++ {

		if seek == s[i] && seekCounter == 1 {
			max++
			i++
			if i == len(s) {
				break
			}
			current = s[i]
			seekCounter = 1
			if current == 'R' {
				seek = 'L'
			} else {
				seek = 'R'
			}

		} else if s[i] == current {
			seekCounter++
		} else if s[i] == seek {
			seekCounter--
		}
	}

	return max
}

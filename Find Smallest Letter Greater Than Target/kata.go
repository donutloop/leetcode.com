package kata

func nextGreatestLetter(letters []byte, target byte) byte {

	lastChar := byte(1)
	for _, c := range letters {
		if target < c {
			lastChar = c
			break
		}
	}

	if lastChar == byte(1) {
		lastChar = letters[0]
	}

	return lastChar
}

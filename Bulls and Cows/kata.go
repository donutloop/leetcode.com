package kata

import "strconv"

func getHint(secret string, guess string) string {
	var bulls int
	var cows int

	chars := make(map[int]int)
	idxs := make(map[int]bool)
	for i, char := range secret {
		chars[i] = int(char)
		charNorm := int(char) * -1
		chars[charNorm] = chars[charNorm] + 1
	}

	for i, char := range guess {
		v, ok := chars[i]
		if ok && v == int(char) {
			charNorm := int(char) * -1
			chars[charNorm] = chars[charNorm] - 1
			if chars[charNorm] == 0 {
				delete(chars, charNorm)
			}
			idxs[i] = true
			bulls++
			continue
		}
	}

	for i, char := range guess {
		if idxs[i] {
			continue
		}
		charNorm := int(char) * -1
		v := chars[charNorm]
		if v > 0 {
			chars[charNorm] = chars[charNorm] - 1
			cows++
		}
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

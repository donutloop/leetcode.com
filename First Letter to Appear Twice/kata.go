package kata

func repeatedCharacter(s string) byte {

	indexLocation := make(map[rune]int)
	for i, char := range s {
		v, ok := indexLocation[char]
		if !ok {
			indexLocation[char] = -1
		} else if v == -1 {
			indexLocation[char] = i
		}
	}

	var minChar rune
	min := -1
	for char, i := range indexLocation {
		if i != -1 && (min == -1 || min > i) {
			min = i
			minChar = char
		}
	}

	return byte(minChar)
}

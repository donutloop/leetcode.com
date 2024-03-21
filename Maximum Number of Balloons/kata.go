package kata

func maxNumberOfBalloons(text string) int {
	chars := make(map[rune]int, 0)
	for _, char := range text {
		if char == 'b' || char == 'a' || char == 'l' || char == 'o' || char == 'n' {
			chars[char]++
		}
	}

	if len(chars) < 5 {
		return 0
	}

	var min = -1
	for char, count := range chars {
		if char == 'l' || char == 'o' {
			count = int(count / 2)
		}
		if min == -1 || min > count {
			min = count
		}
	}
	return min
}

package Single_Row_Keyboard

import "math"

func calculateTime(keyboard string, word string) int {
	keyMap := make(map[rune]int)
	for i, key := range keyboard {
		keyMap[key] = i
	}
	var startPoint int
	var sum int
	for _, char := range word {
		sum += int(math.Abs(float64(keyMap[char] - startPoint)))
		startPoint = keyMap[char]
	}
	return sum
}

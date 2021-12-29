package Check_Whether_Two_Strings_are_Almost_Equivalent

import "math"

func checkAlmostEquivalent(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	wordCounter1 := [26]int{}
	wordCounter2 := [26]int{}
	for i := 0; i < len(word1); i++ {
		wordCounter1[int(word1[i])-97]++
		wordCounter2[int(word2[i])-97]++
	}

	for i := 0; i < len(wordCounter1); i++ {
		diff := math.Abs(float64(wordCounter1[i] - wordCounter2[i]))
		if diff > 3.0 {
			return false
		}
	}

	return true
}

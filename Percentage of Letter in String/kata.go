package Percentage_of_Letter_in_String

import "math"

func percentageLetter(s string, letter byte) int {
	c := 0
	for _, char := range s {
		if rune(letter) == char {
			c++
		}
	}

	if c > 0 {
		return int(math.Floor(float64(c) / float64(len(s)) * 100))
	}

	return c
}

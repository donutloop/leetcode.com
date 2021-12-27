package kata

import "math"

func areOccurrencesEqual(s string) bool {

	chars := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}

	inf := int(math.Inf(1))
	min := inf
	max := inf
	for _, count := range chars {
		if max == inf || count > max {
			max = count
		}
		if min == inf || count < min {
			min = count
		}
	}

	return max == min
}

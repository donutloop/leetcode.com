package kata

import "math"

func MaxPower(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	count := 1

	maxCount := int(math.Inf(-1))
	lastChar := s[0]
	for i := 1; i < len(s); i++ {

		if s[i] == lastChar {
			count++
		} else {
			lastChar = s[i]
			count = 1
		}

		if maxCount == int(math.Inf(-1)) || count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
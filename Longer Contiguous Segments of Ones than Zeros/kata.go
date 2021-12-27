package kata

import "math"

func checkZeroOnes(s string) bool {

	zeroCountMax := int(math.Inf(-1))
	oneCountMax := int(math.Inf(-1))
	var zeroCount int
	var oneCount int

	for i := 0; i < len(s); i++ {
		if zeroCountMax == int(math.Inf(-1)) || zeroCount > zeroCountMax {
			zeroCountMax = zeroCount

		}

		if oneCountMax == int(math.Inf(-1)) || oneCount > oneCountMax {
			oneCountMax = oneCount
		}

		if s[i] == '0' {
			oneCount = 0
			zeroCount++

		} else if s[i] == '1' {
			zeroCount = 0
			oneCount++
		}
		if i == len(s)-1 {

			if zeroCountMax == int(math.Inf(-1)) || zeroCount > zeroCountMax {
				zeroCountMax = zeroCount

			}

			if oneCountMax == int(math.Inf(-1)) || oneCount > oneCountMax {
				oneCountMax = oneCount
			}
		}

	}
	return oneCountMax > zeroCountMax
}

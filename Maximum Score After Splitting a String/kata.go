package Maximum_Score_After_Splitting_a_String

func maxScore(s string) int {

	ones := 0
	for _, char := range s {
		if char != '0' {
			ones++
		}
	}

	max := 0

	leftZeros := 0
	for i, char := range s {
		if char == '0' {
			leftZeros++
		} else {
			ones--
		}

		if ones+leftZeros > max && i != len(s)-1 {
			max = ones + leftZeros
		}
	}

	if leftZeros == len(s) {
		return leftZeros - 1
	}

	return max
}

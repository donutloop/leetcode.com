package kata

func maximumOddBinaryNumber(s string) string {
	var countOnes int
	var countAll int
	for _, digit := range s {
		if digit == '1' {
			countOnes++
		}
		countAll++
	}

	if countOnes == 0 {
		return s
	}

	maximumOddBinary := make([]byte, countAll)
	countOnes--
	maximumOddBinary[countAll-1] = '1'
	for i := 0; i < countAll-1; i++ {
		if countOnes > 0 {
			maximumOddBinary[i] = '1'
			countOnes--
		} else {
			maximumOddBinary[i] = '0'
		}
	}
	return string(maximumOddBinary)
}

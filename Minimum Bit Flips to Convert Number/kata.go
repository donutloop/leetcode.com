package Minimum_Bit_Flips_to_Convert_Number

func minBitFlips(start int, goal int) int {

	maxNum, minNum := max(start, goal), min(start, goal)

	c := maxNum
	i := 0
	k := 0
	for c != 0 {
		c = c / 2

		a := 0
		if maxNum&(1<<i) > 0 {
			a = 1
		}

		b := 0
		if minNum&(1<<i) > 0 {
			b = 1
		}

		if a != b {
			k++
		}

		i++
	}

	return k
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

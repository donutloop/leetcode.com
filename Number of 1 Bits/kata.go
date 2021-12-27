package kata

func hammingWeight(num uint32) int {
	var countSetBits int
	for num > 0 {
		bit := num % 2
		num = num / 2
		if bit == 1 {
			countSetBits++
		}
	}
	return countSetBits
}

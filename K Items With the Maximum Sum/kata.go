package kata

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	var sum int
	if numOnes >= k {
		sum += 1 * k
		return sum
	} else if numOnes > 0 {
		sum += numOnes * 1
		k = k - numOnes
	}
	if numZeros >= k {
		return sum
	} else if numZeros > 0 {
		k = k - numZeros
	}
	if numNegOnes >= k {
		sum += -1 * k
		return sum
	} else if numNegOnes > 0 {
		sum += -1 * numNegOnes
	}
	return sum
}

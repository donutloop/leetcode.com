package kata

func minimumAverageDifference(nums []int) int {
	var sum int
	var count int
	for _, num := range nums {
		sum += num
		count++
	}

	var prefixSum int
	suffixSum := sum
	suffixCount := count
	var min = -1
	var idx = 0
	for i, num := range nums {
		suffixSum = suffixSum - num
		suffixCount--
		prefixSum += num

		var a int
		if suffixSum != 0 {
			a = (suffixSum / suffixCount)
		}

		v := abs((prefixSum / (i + 1)) - a)
		if min == -1 || v < min {
			min = v
			idx = i
		}
	}

	return idx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

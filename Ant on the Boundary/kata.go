func returnToBoundaryCount(nums []int) int {
	var prefixSum int
	var count int
	for _, num := range nums {
		prefixSum += num
		if prefixSum == 0 {
			count++
		}
	}
	return count
}
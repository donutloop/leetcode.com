func maximizeSum(nums []int, k int) int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	var score int
	score += max
	for i := 1; i < k; i++ {
		max = max + 1
		score += max
	}
	return score
}
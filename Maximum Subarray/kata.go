package kata

import "math"

func maxSubArray(nums []int) int {
	var currentSum int
	var globalMax int
	var negativCounter int
	max := math.Inf(-1)
	for _, num := range nums {
		if num < 0 {
			negativCounter++
		}
		if max == math.Inf(-1) || float64(num) > max {
			max = float64(num)
		}
		currentSum = int(math.Max(0, float64(currentSum+num)))
		globalMax = int(math.Max(float64(globalMax), float64(currentSum)))
	}

	if negativCounter == len(nums) {
		return int(max)
	}

	return globalMax
}

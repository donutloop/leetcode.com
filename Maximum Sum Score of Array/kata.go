package Maximum_Sum_Score_of_Array

import "math"

func maximumSumScore(nums []int) int64 {

	var globalMax int64
	for _, n := range nums {
		globalMax += int64(n)
	}

	finalMax := int64(math.MaxInt64) * -1

	rightMax := int64(nums[len(nums)-1])
	r := max(rightMax, globalMax)
	if r > finalMax {
		finalMax = r
	}

	leftMax := int64(nums[0])
	r = max(leftMax, globalMax)
	if r > finalMax {
		finalMax = r
	}

	if len(nums) == 2 {
		return finalMax
	}

	leftMax = int64(0)
	leftMaxGlobal := globalMax
	for i := 0; i < len(nums)-1; i++ {
		leftMax += int64(nums[i])
		leftMaxGlobal -= int64(nums[i])
		r := max(leftMax, leftMaxGlobal)
		if r > finalMax {
			finalMax = r
		}
	}

	rightMax = int64(0)
	rightMaxGlobal := globalMax
	for i := len(nums) - 1; i > 0; i-- {
		rightMax += int64(nums[i])
		rightMaxGlobal -= int64(nums[i])
		r := max(rightMax, rightMaxGlobal)
		if r > finalMax {
			finalMax = r
		}
	}

	return finalMax
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

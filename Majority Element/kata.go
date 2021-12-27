package kata

import (
	"math"
)

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	max := int(math.Ceil(float64(n) / 2))
	counter := make(map[int]int)
	var val int
	for i := 0; i < len(nums); i++ {
		counter[nums[i]] = counter[nums[i]] + 1
		if counter[nums[i]] == max {
			val = nums[i]
			break
		}
	}
	return val
}

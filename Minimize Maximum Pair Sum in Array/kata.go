package kata

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	var max int
	for ; i < len(nums)/2; i++ {
		sum := nums[i] + nums[j]
		if sum > max {
			max = sum
		}
		j--
	}

	return max
}

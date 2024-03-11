package kata

import "sort"

func matrixSum(nums [][]int) int {

	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}

	var sum int
	for i := 0; i < len(nums[0]); i++ {
		var max int
		for j := len(nums) - 1; j >= 0; j-- {
			if max < nums[j][i] {
				max = nums[j][i]
			}
		}
		sum += max
	}

	return sum
}

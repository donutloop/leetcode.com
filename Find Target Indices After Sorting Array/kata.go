package kata

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var indices []int

	j := len(nums) - 1
	for i := 0; i <= len(nums)/2; i++ {
		if nums[i] == target {
			indices = append(indices, i)
		}

		if j == len(nums)/2 {
			continue
		}

		if nums[j] == target {
			indices = append(indices, j)
		}
		j--
	}

	sort.Ints(indices)
	return indices
}

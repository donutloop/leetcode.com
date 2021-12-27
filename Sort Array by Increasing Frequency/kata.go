package kata

import "sort"

func frequencySort(nums []int) []int {

	frequencies := make(map[int]int)
	for _, n := range nums {
		frequencies[n]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if frequencies[nums[i]] < frequencies[nums[j]] {
			return true
		} else if frequencies[nums[i]] == frequencies[nums[j]] {
			return nums[i] > nums[j]
		}
		return false
	})

	return nums
}

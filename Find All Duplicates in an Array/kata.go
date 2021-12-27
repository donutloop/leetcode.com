package kata

import "sort"

func findDuplicates(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var j int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums[j] = nums[i]
			i++
			j++
			continue
		}
	}
	return nums[:j]
}

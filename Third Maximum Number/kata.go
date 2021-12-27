package kata

import (
	"sort"
)

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	seen := map[int]struct{}{}
	count := 0
	for i := 0; i < len(nums); i++ {
		_, ok := seen[nums[i]]
		if ok {
			continue
		}
		seen[nums[i]] = struct{}{}
		count = count + 1
		if count == 3 {
			return nums[i]
		}
	}
	return nums[0]
}

package kata

import (
	"math"
	"sort"
)

func removeElement(nums []int, val int) int {

	var counter int
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = int(math.Inf(1))
			counter++
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	nums = nums[counter:]

	return len(nums)
}

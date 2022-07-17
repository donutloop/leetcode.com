package kata

import "sort"

func findLonely(nums []int) []int {
	sort.Ints(nums)

	var res []int
	for i := 1; i < len(nums); i++ {
		if check(nums, i) {
			continue
		}
		res = append(res, nums[i])
	}

	if !check(nums, 0) {
		res = append(res, nums[0])
	}

	return res
}

func check(nums []int, i int) bool {
	if (i > 0 && nums[i-1] == nums[i]-1) || (i+1 < len(nums) && nums[i+1] == nums[i]+1) || (i > 0 && nums[i-1] == nums[i]) || (i+1 < len(nums) && nums[i+1] == nums[i]) {
		return true
	}
	return false
}

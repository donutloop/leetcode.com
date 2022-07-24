package kata

import "sort"

func numberOfPairs(nums []int) []int {

	sort.Ints(nums)
	var r int
	var p int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			i++
			p++
		} else {
			r++
		}
	}

	if p == 0 {
		r++
	}

	if 2*p+r != len(nums) {
		r++
	}

	return []int{p, r}
}

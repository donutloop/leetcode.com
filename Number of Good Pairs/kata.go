package kata

import "sort"

func numIdenticalPairs(nums []int) int {

	sort.Ints(nums)

	var c int
	var knowNumber *int
	var k int
	for i := 0; i < len(nums); i++ {
		if knowNumber == nil {
			knowNumber = new(int)
			*knowNumber = i
			k++
			continue
		}

		if nums[*knowNumber] != nums[i] || (i == len(nums)-1 && nums[i] == nums[*knowNumber]) {

			if i == len(nums)-1 && nums[i] == nums[*knowNumber] {
				k++
			}

			if k == 1 {
				*knowNumber = i
				k = 1
				continue
			}

			if k == 2 {
				c += 1
			} else {
				c += (k - 1) * (k) / 2
			}

			*knowNumber = i
			k = 0
		}
		k++

	}

	return c
}

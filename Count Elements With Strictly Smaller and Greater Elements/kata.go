package Count_Elements_With_Strictly_Smaller_and_Greater_Elements

import "sort"

func countElements(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	var k int
	var t int
	for i := 2; i < len(nums); i++ {
		if nums[i-2] < nums[i-1] && nums[i-1] < nums[i] {
			k++
		} else if nums[i-2] < nums[i-1] && nums[i-1] == nums[i] {
			t = 2
		} else if nums[i-2] == nums[i-1] && nums[i-1] < nums[i] {
			k += t
			t = 0
		} else if nums[i-2] == nums[i-1] && nums[i-1] == nums[i] {
			if t > 0 {
				t++
			}
		}
	}

	return k
}

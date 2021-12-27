package kata

import "math"

func smallestEqual(nums []int) int {

	index := int(math.Inf(-1))
	for i, n := range nums {
		if i%10 == n {
			if index == int(math.Inf(-1)) || index > i {
				index = i
			}
		}
	}

	if index == int(math.Inf(-1)) {
		return -1
	}

	return index
}

package kata

import "math"

func getMinDistance(nums []int, target int, start int) int {

	secondRound := false
	min := int(math.Inf(-1))
	for i := start; i < len(nums); i++ {

		if nums[i] == target {

			d := int(math.Abs(float64(i) - float64(start)))
			if min == int(math.Inf(-1)) || min > d {
				min = d
			}
		}
		if !secondRound && i+1 == len(nums) {
			i = -1
			secondRound = true
		}
	}

	return min
}

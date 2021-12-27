package kata

import (
	"math"
	"sort"
)

func majorityElement(nums []int) []int {

	sort.Ints(nums)
	currentNumber := nums[0]
	border := math.Floor(float64(len(nums)) / 3)
	count := 1
	var j int
	for i := 1; i < len(nums); i++ {
		if currentNumber == nums[i] {
			count++
		} else {
			if float64(count) > border {
				nums[j] = currentNumber
				j++
			}
			currentNumber = nums[i]
			count = 1
		}
	}

	if count > 0 && float64(count) > border {
		nums[j] = currentNumber
		j++
	}

	return nums[:j]
}

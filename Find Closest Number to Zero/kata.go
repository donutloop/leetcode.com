package Find_Closest_Number_to_Zero

import (
	"math"
	"sort"
)

func findClosestNumber(nums []int) int {

	sort.Ints(nums)

	min := math.MaxInt64
	value := 0
	for _, n := range nums {
		nn := int(math.Abs(float64(n)))
		if nn <= min {
			min = nn
			value = n
		}
	}

	return value
}

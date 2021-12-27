package kata

import (
	"math"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	currentNum := 1
	preNum := math.Inf(-1)
	for _, num := range nums {
		if num > 0 {
			if int(preNum) == num {
				continue
			}
			if currentNum == num {
				preNum = float64(num)
				currentNum++
				continue
			}
			break
		}
	}
	return currentNum
}

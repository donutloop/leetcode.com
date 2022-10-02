package Most_Frequent_Even_Element

import "sort"

func mostFrequentEven(nums []int) int {
	sort.Ints(nums)

	var max int
	var globalMax int
	var currentMaxNumber int
	var currentNumber int
	for i, n := range nums {
		if n%2 == 0 || n == 0 {
			if n != currentNumber {
				if globalMax < max {
					globalMax = max
					currentMaxNumber = nums[i-1]
				}
				max = 0
			}
			currentNumber = n
			max++
		} else {
			if globalMax < max {
				globalMax = max
				currentMaxNumber = nums[i-1]
			}
			max = 0
		}
	}

	if globalMax < max {
		currentMaxNumber = nums[len(nums)-1]
		globalMax = max
	}
	if globalMax == 0 {
		currentMaxNumber = -1
	}

	return currentMaxNumber
}

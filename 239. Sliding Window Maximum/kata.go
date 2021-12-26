package _39__Sliding_Window_Maximum

import "math"

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 1 {
		return nums
	}

	var maxNums []int
	max := findMax(nums, 0, k-1)
	maxNums = append(maxNums, max.n)

	for i := k; i < len(nums); i++ {

		j := i - k + 1
		if max.i < j {
			max = findMax(nums, j, i)
		}

		if max.n > nums[i] {
			maxNums = append(maxNums, max.n)
		} else {
			maxNums = append(maxNums, nums[i])
			max = maxNumber{n: nums[i], i: i}
		}
	}

	return maxNums
}

type maxNumber struct {
	i int
	n int
}

func findMax(nums []int, i, k int) maxNumber {
	max := int(math.Inf(-1))
	var j int
	for ; i <= k; i++ {
		if nums[i] > max {
			max = nums[i]
			j = i
		}
	}
	return maxNumber{i: j, n: max}
}

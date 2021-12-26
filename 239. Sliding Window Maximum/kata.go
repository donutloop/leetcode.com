package _39__Sliding_Window_Maximum

import "math"

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 1 {
		return nums
	}

	l := 0
	max := findMax(nums, 0, k-1)
	nums[l] = max.n

	for i := k; i < len(nums); i++ {
		l++
		j := i - k + 1
		if max.i < j {
			max = findMax(nums, j, i)
		}

		if max.n > nums[i] {
			nums[l] = max.n
		} else {
			nums[l] = nums[i]
			max = maxNumber{n: nums[i], i: i}
		}
	}

	return nums[:l+1]
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

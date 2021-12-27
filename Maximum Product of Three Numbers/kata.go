package kata

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	var sumA, sumB = 1, 1

	for _, n := range nums[len(nums)-3:] {
		sumB = n * sumB
	}

	for _, n := range nums[:3] {
		sumA = n * sumA
	}

	if sumA < sumB {
		sumA = sumB
	}
	sumB = 1

	for _, n := range nums[:2] {
		sumB = n * sumB
	}
	sumB = sumB * nums[len(nums)-1]

	if sumA > sumB {
		return sumA
	}

	return sumB
}

package Minimum_Sum_of_Four_Digit_Number_After_Splitting_Digits

import "sort"

func minimumSum(num int) int {

	var nums []int
	for {
		r := num % 10
		num = num / 10
		nums = append(nums, r)
		if num == 0 {
			break
		}
	}

	sort.Ints(nums)

	uB, uA := 1, 1
	var a, b int

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			continue
		}

		if (uA*nums[i])+a > (uB*nums[i])+b {
			b = (uB * nums[i]) + b
			uB = uB * 10
		} else {
			a = (uA * nums[i]) + a
			uA = uA * 10
		}

	}

	return a + b
}

package Minimum_Operations_to_Make_the_Array_Increasing

func minOperations(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	var operations int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			continue
		}
		n := nums[i-1] - nums[i] + 1
		operations += n
		nums[i] += n
	}

	return operations
}

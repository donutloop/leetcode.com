package kata

func buildArray(nums []int) []int {

	permNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		permNums[i] = nums[nums[i]]
	}

	return permNums
}

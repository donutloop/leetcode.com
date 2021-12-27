package kata

func runningSum(nums []int) []int {
	sumNums := make([]int, len(nums))
	subTotal := 0
	for i := 0; i < len(nums); i++ {
		subTotal += nums[i]
		sumNums[i] = subTotal
	}

	return sumNums
}

package kata

func pivotIndex(nums []int) int {
	var sum int
	for _, n := range nums {
		sum = sum + n
	}

	var leftSum int
	last := -1
	for i := len(nums) - 1; i >= 0; i-- {
		r := sum - nums[i]
		l := leftSum
		if r == l {
			last = i
		}
		leftSum = leftSum + nums[i]
		sum = sum - nums[i]
	}
	return last
}

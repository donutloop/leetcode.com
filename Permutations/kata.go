package kata

func permute(nums []int) [][]int {
	permutations := make([][]int, 0)
	perm(nums, &permutations, 0)
	return permutations
}

func perm(nums []int, permutations *[][]int, i int) {
	if i > len(nums) {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		*permutations = append(*permutations, numsCopy)
		return
	}
	perm(nums, permutations, i+1)
	for j := i + 1; j < len(nums); j++ {
		// swap i and j value
		nums[i], nums[j] = nums[j], nums[i]
		perm(nums, permutations, i+1)
		// revert swap between i and j values
		nums[i], nums[j] = nums[j], nums[i]
	}
}

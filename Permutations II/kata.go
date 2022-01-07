package Permutations_II

func permuteUnique(nums []int) [][]int {
	permutations := make([][]int, 0)
	duplicates := make(map[int]bool)
	perm(nums, &permutations, 0, duplicates)
	return permutations
}

func perm(nums []int, permutations *[][]int, i int, duplicates map[int]bool) {
	if i > len(nums) {
		n := 1
		var sum int
		for i := 0; i < len(nums); i++ {
			sum += nums[i] * n
			n = n * 10
		}

		if duplicates[sum] {
			return
		}

		duplicates[sum] = true

		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		*permutations = append(*permutations, numsCopy)
		return
	}
	perm(nums, permutations, i+1, duplicates)
	for j := i + 1; j < len(nums); j++ {
		nums[i], nums[j] = nums[j], nums[i]
		perm(nums, permutations, i+1, duplicates)
		nums[i], nums[j] = nums[j], nums[i]
	}
}

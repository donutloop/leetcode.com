package kata

func missingNumber(nums []int) int {
	count := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		count[nums[i]] = count[nums[i]] + 1
	}
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			return i
		}
	}
	return -1
}

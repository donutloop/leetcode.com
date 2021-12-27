package kata

func containsDuplicate(nums []int) bool {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]] = counter[nums[i]] + 1
		if counter[nums[i]] > 1 {
			return true
		}
	}
	return false
}

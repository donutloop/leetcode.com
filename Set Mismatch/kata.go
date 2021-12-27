package kata

func findErrorNums(nums []int) []int {
	counter := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		counter[nums[i]-1] = counter[nums[i]-1] + 1
	}
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if counter[i] == 0 {
			res[1] = i + 1
		} else if counter[i] > 1 {
			res[0] = i + 1
		}
	}
	return res
}

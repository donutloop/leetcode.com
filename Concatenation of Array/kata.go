package kata

func getConcatenation(nums []int) []int {

	numsDouble := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		numsDouble[i] = nums[i]
		numsDouble[i+len(nums)] = nums[i]
	}

	return numsDouble
}

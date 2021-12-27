package kata

func decompressRLElist(nums []int) []int {
	values := make([]int, 0, len(nums)/2)
	for i := 0; i < len(nums); i = i + 2 {
		for j := 0; j < nums[i]; j++ {
			values = append(values, nums[i+1])
		}
	}
	return values
}

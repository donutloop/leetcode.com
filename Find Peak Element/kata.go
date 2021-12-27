package kata

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxPeak := -1
	idx := 0
	for i := 0; i < len(nums); i++ {
		if maxPeak == -1 || maxPeak < nums[i] {
			maxPeak = nums[i]
			idx = i
		}
	}
	return idx
}

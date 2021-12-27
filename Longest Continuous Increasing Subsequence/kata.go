package kata

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	a := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			a++
			continue
		}
		if a > l {
			l = a
		}
		a = 0
	}
	if a > l {
		l = a
	}
	return l + 1
}

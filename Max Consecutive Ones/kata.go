package kata

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var c int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if c > max {
				max = c
			}
			c = 0
			continue
		}
		c++
	}
	if c > max {
		max = c
	}
	return max
}

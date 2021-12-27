package kata

func shuffle(nums []int, n int) []int {
	shuffled := make([]int, len(nums))
	toggle := true
	j := 0
	for i := 0; i < len(nums); i++ {
		if toggle {
			shuffled[i] = nums[j]
		} else {
			shuffled[i] = nums[j+n]
			j++
		}
		toggle = !toggle
	}
	return shuffled
}

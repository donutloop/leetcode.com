package kata

func findDisappearedNumbers(nums []int) []int {
	b := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		b[nums[i]] = 1
	}
	c := make([]int, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			c = append(c, i)
		}
	}
	return c[1:]
}

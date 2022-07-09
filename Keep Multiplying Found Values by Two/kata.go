package Keep_Multiplying_Found_Values_by_Two

func findFinalValue(nums []int, original int) int {

	var i int
	for {
		i = i % len(nums)
		n := nums[i]
		if n == original {
			i = 0
			original = original * 2
			continue
		}
		i++
		if i == len(nums) {
			break
		}
	}
	return original
}

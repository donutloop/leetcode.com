package kata

func hasTrailingZeros(nums []int) bool {
	var count int
	for _, num := range nums {
		if num%2 == 0 {
			count++
			if count == 2 {
				break
			}
		}
	}
	return count == 2
}

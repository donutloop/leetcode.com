package kata

func isPossibleToSplit(nums []int) bool {
	frequency := make(map[int]int, 0)
	for _, num := range nums {
		frequency[num]++
		if frequency[num] > 2 {
			return false
		}
	}
	return true
}

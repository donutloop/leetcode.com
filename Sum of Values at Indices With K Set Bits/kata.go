package kata

func sumIndicesWithKSetBits(nums []int, k int) int {
	var sum int
outerLoop:
	for i, num := range nums {
		var count int
		for i > 0 {
			count += i & 1
			if count > k {
				continue outerLoop
			}
			i >>= 1
		}
		if count == k {
			sum += num
		}
	}
	return sum
}

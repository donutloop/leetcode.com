package kata

func countBits(n int) []int {
	o := make([]int, n+1)
	for k := 0; k <= n; k++ {
		var num = k
		var count int
		for num > 0 {
			count += num & 1
			num >>= 1
		}
		o[k] = count
	}
	return o
}

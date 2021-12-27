package kata

func xorOperation(n int, start int) int {
	var k int
	for i := 0; i < n; i++ {
		k ^= start + 2*i
	}
	return k
}

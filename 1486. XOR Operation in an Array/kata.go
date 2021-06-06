package _486__XOR_Operation_in_an_Array

func xorOperation(n int, start int) int {
	var k int
	for i := 0; i < n; i++ {
		if i > 0 {
			k = k ^ (start + 2*i)
		} else {
			k = start + 2*i
		}
	}
	return k
}

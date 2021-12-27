package kata

func hasAlternatingBits(n int) bool {
	lastBit := n % 2
	n = n >> 1
	for n > 0 {
		b := n % 2
		n = n >> 1
		if lastBit == b {
			return false
		}
		lastBit = b
	}
	return true
}

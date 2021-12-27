package kata

func countOdds(low int, high int) int {
	var count int
	for ; low < high+1; low++ {
		if low%2 == 1 {
			count++
		}
	}
	return count
}

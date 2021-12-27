package kata

func isMonotonic(A []int) bool {
	var dec, inc int
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			inc++
		} else if A[i-1] > A[i] {
			dec++
		}
	}
	if (dec == 0 && inc != 0) || (dec != 0 && inc == 0) || (dec == 0 && inc == 0) {
		return true
	}

	return false
}

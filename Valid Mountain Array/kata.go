package kata

func validMountainArray(A []int) bool {
	if len(A) < 2 {
		return false
	}
	if A[0] >= A[1] || A[len(A)-2] < A[len(A)-1] {
		return false
	}
	var down bool
	for i := 1; i < len(A); i++ {
		if A[i-1] > A[i] {
			down = true
		}
		if (A[i-1] < A[i] && down) || A[i-1] == A[i] {
			return false
		}
	}
	return true
}

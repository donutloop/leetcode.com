package kata

func repeatedNTimes(A []int) (n int) {
	set := make(map[int]int)
	for i := 0; i < len(A); i++ {
		set[A[i]]++
		if set[A[i]] > 1 {
			n = A[i]
			return
		}
	}
	return
}

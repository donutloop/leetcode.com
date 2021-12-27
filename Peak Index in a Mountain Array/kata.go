package kata

func peakIndexInMountainArray(A []int) int {
	max := -1
	idx := -1
	for i := 0; i < len(A); i++ {
		if max == -1 || A[i] > max {
			max = A[i]
			idx = i
		}
	}
	return idx
}

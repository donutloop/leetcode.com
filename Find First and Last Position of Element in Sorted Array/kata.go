package kata

func searchRange(nums []int, target int) []int {

	idx := make([]int, 2)
	idx[0] = -1
	idx[1] = -1

	first := true
	for i, n := range nums {
		if n == target {
			if first {
				idx[0] = i
				first = false
			} else {
				idx[1] = i
			}
		}
	}

	if idx[1] == -1 {
		idx[1] = idx[0]
	}

	return idx
}

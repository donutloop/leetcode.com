package kata

func resultArray(nums []int) []int {
	var a, b = make([]int, 0), make([]int, 0)

	a = append(a, nums[0])
	b = append(b, nums[1])

	for i := 2; i < len(nums); i++ {
		if a[len(a)-1] > b[len(b)-1] {
			a = append(a, nums[i])
		} else {
			b = append(b, nums[i])
		}
	}

	for _, ib := range b {
		a = append(a, ib)
	}

	return a
}

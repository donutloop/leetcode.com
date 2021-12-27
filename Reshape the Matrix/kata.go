package kata

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nums == nil {
		return nums
	}
	if len(nums[0])*len(nums) != r*c {
		return nums
	}
	res := make([][]int, 0)
	n := c
	i := 0
	res = append(res, make([]int, 0))
	for _, row := range nums {
		for _, col := range row {
			if n == 0 {
				n = c
				i = i + 1
				res = append(res, make([]int, 0))
			}
			n = n - 1
			res[i] = append(res[i], col)
		}
	}
	return res
}

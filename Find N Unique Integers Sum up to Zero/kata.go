package kata

func sumZero(n int) []int {
	nums := make([]int, 0)
	if n%2 == 1 {
		nums = append(nums, 0)
		n = n - 1
	}
	for i := 1; i < n; i = i + 2 {
		nums = append(nums, i, -i)
	}
	return nums
}

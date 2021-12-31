package get_Maximum_in_Generated_Array

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}

	nums := make([]int, n+1)
	var l int
	nums[l] = 0
	l++
	nums[l] = 1
	max := -1
	for ; l < (n+1)/2; l++ {
		i := 2 * l
		nums[i] = nums[l]
		if max < nums[i] {
			max = nums[i]
		}
		nums[i+1] = nums[l] + nums[l+1]
		if max < nums[i+1] {
			max = nums[i]
		}
	}

	for ; l <= n; l++ {
		if max < nums[l] {
			max = nums[l]
		}
	}

	return max
}

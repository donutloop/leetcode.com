package K_Radius_Subarray_Averages

func getAverages(nums []int, k int) []int {

	if k == 0 {
		return nums
	}

	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		out[i] = -1
	}

	if k > len(nums) || (k+k) >= len(nums) {
		return out
	}

	var sum int
	var c int
	for i := 0; i < k+k+1; i++ {
		sum += nums[i]
		c++
	}

	avg := sum / c
	out[k] = avg

	for i := k + 1; i < len(nums); i++ {
		if i+k >= len(nums) {
			break
		}

		sum += nums[i+k]
		sum -= nums[i-k-1]
		out[i] = sum / c
	}

	return out
}

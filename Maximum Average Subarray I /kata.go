package Maximum_Average_Subarray_I_

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := range nums {
		if i == k {
			break
		}
		sum = sum + nums[i]
	}

	max := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k]
		sum = sum + nums[i]
		if sum > max {
			max = sum
		}
	}

	return float64(max) / float64(k)
}

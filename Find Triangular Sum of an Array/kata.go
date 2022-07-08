package Find_Triangular_Sum_of_an_Array

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	i := len(nums) - 1
	j := len(nums) - 2
	k := 1
	for {
		nums = append(nums, (nums[k-1]+nums[k])%10)
		k++
		i--
		if i == 0 {
			k++
			i += j
			j = j - 1
		}
		if i == 0 {
			break
		}
	}

	return nums[len(nums)-1]
}

package kata

func containsNearbyDuplicate(nums []int, k int) bool {

	i := 0
	j := i + 1
	for i < len(nums) {

		if j == len(nums) {
			i++
			j = i + 1
			continue
		}
		if nums[i] != nums[j] {
			j++
			continue
		}

		if (j - i) <= k {
			return true
		}
		i++
		j = i + 1
	}

	return false
}

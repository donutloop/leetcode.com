package kata

func createTargetArray(nums []int, index []int) []int {
	if len(nums) != len(index) {
		panic("index and nums len is unequal")
	}

	target := make([]int, len(nums))
	for i := range target {
		target[i] = -1
	}

	for i := range nums {
		if target[index[i]] > -1 {
			val := target[index[i]]
			target[index[i]] = nums[i]

			j := index[i] + 1
			for j < len(target) {
				jVal := target[j]
				target[j] = val
				val = jVal
				j++
			}
			continue
		} else {
			target[index[i]] = nums[i]
		}
	}
	return target
}

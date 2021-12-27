package kata

func dominantIndex(nums []int) int {
	var index int
	var max int
	var i int
	for i = 0; i < len(nums); i++ {
		if (max * 2) <= nums[i] {
			max = nums[i]
			index = i
		}
	}
	for i = 0; i < len(nums); i++ {
		if max != nums[i] && max < (2*nums[i]) {
			return -1
		}
	}
	return index
}

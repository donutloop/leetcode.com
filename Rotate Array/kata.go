package kata

func rotate(nums []int, k int) {
	for ; k > 0; k-- {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = tmp
	}
}

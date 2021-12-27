package _36__Single_Number

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	x := nums[0]
	for _, n := range nums[1:] {
		x = n ^ x
	}
	return x
}

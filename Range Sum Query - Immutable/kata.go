package kata

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if j+1 > len(this.nums) || i > len(this.nums) || i < 0 || j < 0 || i > j {
		return 0
	}

	var sum int
	for _, n := range this.nums[i : j+1] {
		sum += n
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

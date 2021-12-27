package kata

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) Update(i int, val int) {
	if i > len(this.nums) {
		return
	}

	this.nums[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	var sum int
	for _, n := range this.nums[i : j+1] {
		sum += n
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

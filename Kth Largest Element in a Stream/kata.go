package kata

import "sort"

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	sort.Slice(this.nums, func(i, j int) bool {
		return this.nums[i] < this.nums[j]
	})
	return this.nums[len(this.nums)-this.k]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

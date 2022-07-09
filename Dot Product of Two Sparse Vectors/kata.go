package Dot_Product_of_Two_Sparse_Vectors

type SparseVector struct {
	Values map[int]int
}

func Constructor(nums []int) SparseVector {
	v := make(map[int]int)
	for i, n := range nums {
		v[i] = n
	}
	return SparseVector{
		Values: v,
	}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	var sum int
	for i, num := range vec.Values {
		sum += this.Values[i] * num
	}
	return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */

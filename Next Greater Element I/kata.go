package kata

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		set[nums2[i]] = i
	}

	n := make([]int, len(nums1))
	o := 0
	for i := 0; i < len(nums1); i++ {
		j := set[nums1[i]]
		found := false
		for k := j + 1; k < len(nums2); k++ {
			if nums2[j] < nums2[k] {
				n[o] = nums2[k]
				found = true
				break
			}
		}
		if !found {
			n[o] = -1
		}
		o++

	}
	return n
}

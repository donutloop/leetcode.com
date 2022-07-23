package kata

func anagramMappings(nums1 []int, nums2 []int) []int {

	indexLocation := make(map[int]int)
	for i, n := range nums2 {
		indexLocation[n] = i
	}

	mapping := make([]int, len(nums1))
	for i, n := range nums1 {
		mapping[i] = indexLocation[n]
	}

	return mapping
}

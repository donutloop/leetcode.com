package kata

func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		seen[nums1[i]] = 1
	}
	inter := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		counter, ok := seen[nums2[i]]
		if ok && counter == 1 {
			seen[nums2[i]] = seen[nums2[i]] + 1
			inter = append(inter, nums2[i])
		}
	}
	return inter
}

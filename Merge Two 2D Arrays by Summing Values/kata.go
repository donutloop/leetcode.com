package kata

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	merged := make([][]int, 0)

	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i][0] == nums2[j][0] {
				nums1[i][1] += nums2[j][1]
				merged = append(merged, nums1[i])
				i++
				j++
			} else if nums1[i][0] < nums2[j][0] {
				merged = append(merged, nums1[i])
				i++
			} else {
				merged = append(merged, nums2[j])
				j++
			}
		} else if i < len(nums1) {
			merged = append(merged, nums1[i])
			i++
		} else if j < len(nums2) {
			merged = append(merged, nums2[j])
			j++
		}
	}

	return merged
}

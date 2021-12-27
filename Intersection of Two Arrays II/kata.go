package kata

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	intersection := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		intersection[nums1[i]] = intersection[nums1[i]] + 1
	}
	y := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		c, ok := intersection[nums2[i]]
		if ok && c > 0 {
			intersection[nums2[i]] = intersection[nums2[i]] - 1
			y = append(y, nums2[i])
		}
	}
	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})
	return y
}

package Intersection_of_Multiple_Arrays

import "sort"

func intersection(nums [][]int) []int {
	set := make(map[int]int)
	for _, vec := range nums {
		for _, n := range vec {
			set[n]++
		}
	}

	vec := make([]int, 0)
	for n, c := range set {
		if c == len(nums) {
			vec = append(vec, n)
		}
	}

	if len(vec) != 0 {
		sort.Ints(vec)
	}

	return vec
}

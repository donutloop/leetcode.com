package Assign_Cookies

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}

	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	var j int
	var c int
	for i := 0; i < len(g); {

		if s[j] >= g[i] {
			c++
			i++
			j++
			if j == len(s) {
				break
			}
		} else {
			i++
		}
	}
	return c
}

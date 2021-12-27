package kata

import "sort"

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	sort.Ints(arr)
	sort.Ints(target)
	for i, elem := range arr {
		if target[i] != elem {
			return false
		}
	}

	return true
}

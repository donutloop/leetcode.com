package kata

import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	if len(arr) == 1 {
		return []int{1}
	}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	set := make(map[int]int)
	sort.Ints(tmp)
	r := 1

	for _, elm := range tmp {
		_, ok := set[elm]
		if !ok {
			set[elm] = r
			r++
		}
	}
	for i := range arr {
		arr[i] = set[arr[i]]
	}
	return arr
}

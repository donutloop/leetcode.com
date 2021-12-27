package kata

import "sort"

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		counter[arr[i]] = counter[arr[i]] + 1
	}

	values := make([]int, len(counter))
	var i int
	for _, c := range counter {
		values[i] = c
		i++
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	last := -1
	for _, a := range values {
		if a == last {
			return false
		}
		last = a
	}
	return true
}

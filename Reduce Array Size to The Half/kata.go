package kata

import "sort"

func minSetSize(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	if len(arr) == 2 {
		return 1
	}

	counter := make(map[int]int)
	for i := range arr {
		counter[arr[i]]++
	}
	if len(counter) == 1 {
		return 1
	}

	countes := make([]int, len(counter))
	var i int
	for _, c := range counter {
		countes[i] = c
		i++
	}

	sort.Slice(countes, func(i, j int) bool {
		return countes[i] > countes[j]
	})

	half := len(arr) / 2
	var sum int
	var many int
	for _, c := range countes {
		sum += c
		many++
		if sum >= half {
			break
		}
	}

	return many
}

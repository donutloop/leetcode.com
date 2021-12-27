package kata

import "sort"

func customSortString(S string, T string) string {
	order := 1
	orders := make(map[rune]int)
	for _, char := range S {
		orders[char] = order
		order++
	}

	K := []rune(T)
	sort.Slice(K, func(i, j int) bool {
		return orders[K[i]] < orders[K[j]]
	})
	return string(K)
}

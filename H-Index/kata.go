package H_Index

import "sort"

func hIndex(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= i+1 {
			return i + 1
		}
	}

	return 0
}

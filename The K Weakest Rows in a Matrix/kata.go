package kata

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	pair := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		var sum int
		for j := 0; j < len(mat[i]); j++ {
			sum += mat[i][j]
		}
		pair[i] = []int{sum, i}
	}
	sort.Slice(pair, func(i, j int) bool {
		if pair[i][0] == pair[j][0] {
			return pair[i][1] < pair[j][1]
		}
		return pair[i][0] < pair[j][0]
	})
	idx := make([]int, k)
	for i := 0; i < k; i++ {
		idx[i] = pair[i][1]
	}
	return idx
}

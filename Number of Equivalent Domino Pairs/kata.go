package Number_of_Equivalent_Domino_Pairs

import "math/big"

func numEquivDominoPairs(dominoes [][]int) int {

	var matrix [100]int

	var c int
	for _, d := range dominoes {
		n := hash(d)
		matrix[n]++
	}

	for _, d := range matrix {
		if d >= 2 {
			c += Binomial(d)
		}
	}

	return c
}

func hash(d []int) int {
	return 10*max(d[0], d[1]) + min(d[0], d[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	cache = make(map[int]int)
}

var cache map[int]int

func Binomial(n int) int {
	if _, ok := cache[n]; ok {
		return cache[n]
	}

	k := new(big.Int)
	return int(k.Binomial(int64(n), 2).Int64())
}

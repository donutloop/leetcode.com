package kata

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func equalPairs(grid [][]int) int {

	type pair struct {
		A, B int
	}

	matcher := make(map[string]pair)
	for i := 0; i < len(grid); i++ {
		k := make([]byte, 0)
		for j := 0; j < len(grid[i]); j++ {
			k = append(k, []byte(strconv.Itoa(grid[i][j]))...)
			k = append(k, byte(','))
		}

		h := md5.Sum(k)
		hs := hex.EncodeToString(h[:])
		a := matcher[hs]
		a.A++
		matcher[hs] = a
	}

	var c int
	for i := 0; i < len(grid[0]); i++ {
		k := make([]byte, 0)
		for j := 0; j < len(grid); j++ {
			k = append(k, []byte(strconv.Itoa(grid[j][i]))...)
			k = append(k, byte(','))
		}

		h := md5.Sum(k)
		hs := hex.EncodeToString(h[:])
		b := matcher[hs]
		b.B++
		matcher[hs] = b
	}

	for _, pair := range matcher {
		c += pair.A * pair.B
	}

	return c
}

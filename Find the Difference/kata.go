package kata

import "sort"

func findTheDifference(s string, t string) byte {

	bs := []byte(s)

	bt := []byte(t)

	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	sort.Slice(bt, func(i, j int) bool {
		return bt[i] < bt[j]
	})

	var max int
	if len(bs) < len(bt) {
		max = len(bs)
	} else {
		max = len(bt)
	}

	for i := 0; i < max; i++ {
		if bs[i] != bt[i] {
			return bt[i]
		}
	}

	return bt[len(t)-1]
}

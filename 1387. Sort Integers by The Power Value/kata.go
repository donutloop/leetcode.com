package _387__Sort_Integers_by_The_Power_Value

import "sort"

func getKth(lo int, hi int, k int) int {

	type powerValue struct {
		n, p int
	}

	ns := make([]powerValue, hi-lo+1)

	var i int
	for x := lo; x <= hi; x++ {
		y := x
		count := 0
		for y != 1 {
			if y%2 == 0 {
				y = y / 2

			} else if y%2 == 1 {
				y = 3*y + 1

			}
			count++
		}

		ns[i] = powerValue{n: x, p: count}
		i++
	}

	sort.Slice(ns, func(i, j int) bool {
		if ns[i].p == ns[j].p {
			return ns[i].n < ns[j].n
		}
		return ns[i].p < ns[j].p
	})

	return ns[k-1].n
}

package kata

import "math"

func shortestToChar(S string, C byte) []int {
	Cs := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			Cs = append(Cs, i)
		}
	}

	allDis := make([]int, 0)
	for i := 0; i < len(S); i++ {
		min := -1
		for j := 0; j < len(Cs); j++ {
			dis := math.Abs(float64(i - Cs[j]))
			if min == -1 || int(dis) < min {
				min = int(dis)
			}
			if min == 0 {
				break
			}
		}
		allDis = append(allDis, min)
	}

	return allDis
}

package kata

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}

	sort.Ints(arr)
	globalDiff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if globalDiff != diff {
			return false
		}
	}

	return true
}

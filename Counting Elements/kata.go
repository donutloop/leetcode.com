package kata

import "math"

func countElements(arr []int) int {
	init := int(math.Inf(-1))

	max := init
	min := init
	set := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		set[arr[i]] += 1

		if max == init || arr[i] > max {
			max = arr[i]
		}

		if min == init || arr[i] < min {
			min = arr[i]
		}
	}

	lastElement := min
	globalMax := 0
	first := true
	for i := min + 1; i < max+1; i++ {
		c, ok := set[i]
		if !ok {
			lastElement = lastElement + 1
			first = false
			continue
		}

		if lastElement+1 == i {
			nextCounter := set[i+1]
			if nextCounter > 0 {
				globalMax += c
			}
			if first {
				globalMax += set[min]
			}
		}
		first = false
		lastElement = lastElement + 1
	}
	return globalMax
}

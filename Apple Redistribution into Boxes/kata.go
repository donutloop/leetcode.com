package kata

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	var applesCount int
	for i := range apple {
		applesCount += apple[i]
	}

	var sum, count int
	for i := range capacity {
		sum += capacity[i]
		count++
		if sum >= applesCount {
			break
		}
	}

	return count
}

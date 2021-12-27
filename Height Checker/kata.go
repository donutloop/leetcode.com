package kata

import "sort"

func heightChecker(heights []int) int {
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)
	sort.Ints(sortedHeights)
	var count int
	for i := range heights {
		if heights[i] != sortedHeights[i] {
			count++
		}
	}
	return count
}

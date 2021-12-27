package kata

import "sort"

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	for len(stones) != 0 && len(stones) != 1 {
		if !sort.IntsAreSorted(stones) {
			sort.Ints(stones)
		}
		diff := stones[len(stones)-1] - stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if diff > 0 {
			stones = append(stones, diff)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

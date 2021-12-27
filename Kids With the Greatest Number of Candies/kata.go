package kata

import "math"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := int(math.Inf(1))
	for _, count := range candies {
		if max == int(math.Inf(1)) || count > max {
			max = count
		}
	}
	facts := make([]bool, len(candies))
	for i, count := range candies {
		if (count + extraCandies) >= max {
			facts[i] = true
		}
	}

	return facts
}

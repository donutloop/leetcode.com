package kata

func largeGroupPositions(S string) [][]int {
	lastSeenChar := S[0]
	firstIdx := 0
	c := 1

	groups := make([][]int, 0)
	for i := 1; i < len(S); i++ {
		if i == len(S)-1 {
			k := i - 1
			if lastSeenChar == S[i] {
				k = i
				c++
			}
			if c >= 3 {
				groups = append(groups, []int{firstIdx, k})
			}
		} else if lastSeenChar != S[i] {
			if c >= 3 {
				groups = append(groups, []int{firstIdx, i - 1})
			}
			firstIdx = i
			lastSeenChar = S[i]
			c = 1
			continue
		} else {
			c++
		}
	}
	return groups
}

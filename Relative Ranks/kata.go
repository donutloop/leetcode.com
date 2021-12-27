package kata

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	scores := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})
	res := make([]string, len(nums))
	set := make(map[int]string, len(nums))
	for i, n := range tmp {
		if len(scores) > 0 {
			set[n] = scores[0]
			scores = scores[1:]
			continue
		}
		set[n] = strconv.Itoa(i + 1)
	}
	for i, n := range nums {
		res[i] = set[n]
	}
	return res
}

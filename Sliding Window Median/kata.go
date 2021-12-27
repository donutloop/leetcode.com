package kata

import "sort"

func medianSlidingWindow(nums []int, k int) []float64 {
	if len(nums) == 1 {
		return []float64{float64(nums[0])}
	}
	medians := make([]float64, 0)
	var j int
	for i := k; i <= len(nums); i++ {
		window := make([]int, k)
		copy(window, nums[j:i])
		sort.Slice(window, func(m, k int) bool {
			return window[m] < window[k]
		})
		if len(window)%2 == 1 {
			medians = append(medians, float64(window[len(window)/2]))
		} else {
			medians = append(medians, (float64(window[len(window)/2-1])+float64(window[len(window)/2]))/2)
		}
		j++
	}
	return medians
}

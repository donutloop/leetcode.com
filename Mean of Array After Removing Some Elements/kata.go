package kata

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	p := int(float64(len(arr)) / float64(100) * float64(5))
	var mean float64
	for _, n := range arr[p : len(arr)-p] {
		mean = mean + float64(n)
	}

	return mean / float64(len(arr[p:len(arr)-p]))
}

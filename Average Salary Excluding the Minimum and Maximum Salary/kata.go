package kata

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)
	salary = salary[1 : len(salary)-1]
	var average int
	for _, n := range salary {
		average = average + n
	}
	return float64(average) / float64(len(salary))
}

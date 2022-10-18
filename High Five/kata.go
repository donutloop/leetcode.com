package High_Five

import "sort"

func highFive(items [][]int) [][]int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] > items[j][1]
		}
		return items[i][0] < items[j][0]
	})

	currentNumber := items[0][0]
	counter := 1
	sum := items[0][1]
	res := make([][]int, 0)
	for i := 1; i < len(items); i++ {
		if counter == 5 && currentNumber != items[i][0] {
			res = append(res, []int{currentNumber, sum / 5})
			currentNumber = items[i][0]
			sum = items[i][1]
			counter = 1
			continue
		} else if counter == 5 {
			continue
		}
		sum += items[i][1]
		counter++
	}

	if counter == 5 {
		res = append(res, []int{currentNumber, sum / 5})
	}

	return res
}

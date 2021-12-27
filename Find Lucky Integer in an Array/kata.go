package kata

import "sort"

func findLucky(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var maxNumber int
	lastNumber := arr[0]
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == lastNumber {
			count++
		} else {
			if count == lastNumber && lastNumber > maxNumber {
				maxNumber = lastNumber
			}
			lastNumber = arr[i]
			count = 1
		}
	}

	if count > 1 {
		if count == lastNumber && lastNumber > maxNumber {
			maxNumber = lastNumber
		}
	}

	if maxNumber == 0 {
		return -1
	}

	return maxNumber
}

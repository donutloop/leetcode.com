package kata

func findKthPositive(arr []int, k int) int {

	lastN := 0
	missingNumbers := 0
	for i := 0; i < len(arr); i++ {

		diff := (arr[i]) - lastN
		if diff > 1 {
			missingNumbers++
			lastN++
			i--
			if missingNumbers == k {
				return lastN
			}

			continue
		}

		lastN = arr[i]
	}

	return len(arr) + k
}

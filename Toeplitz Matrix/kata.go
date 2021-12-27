package kata

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix[0]); i++ {
		j := 0
		k := i
		lastValue := matrix[0][k]
		for {

			if len(matrix) <= j {
				break
			}

			if len(matrix[0]) <= k {
				break
			}

			if lastValue != matrix[j][k] {
				return false
			}

			lastValue = matrix[j][k]

			j = j + 1
			k++
		}
	}

	for i := len(matrix) - 1; i >= 1; i-- {
		j := i
		k := 0
		lastValue := matrix[j][k]
		for {

			if len(matrix) <= j {
				break
			}

			if len(matrix[0]) <= k {
				break
			}

			if lastValue != matrix[j][k] {
				return false
			}

			lastValue = matrix[j][k]

			j = j + 1
			k++
		}
	}

	return true
}

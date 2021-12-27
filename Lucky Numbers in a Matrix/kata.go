package kata

func luckyNumbers(matrix [][]int) []int {

	type point struct {
		i, j int
	}

	allMins := make([]point, 0)
	for i := 0; i < len(matrix); i++ {
		currentMin := -1
		var p point
		for j := 0; j < len(matrix[i]); j++ {
			if currentMin == -1 || currentMin > matrix[i][j] {
				currentMin = matrix[i][j]
				p = point{i: i, j: j}
			}
		}
		allMins = append(allMins, p)
	}

	validValues := make([]int, 0)
outerLoop:
	for _, point := range allMins {
		for i := 0; i < len(matrix); i++ {
			if matrix[point.i][point.j] < matrix[i][point.j] {
				continue outerLoop
			}
		}
		validValues = append(validValues, matrix[point.i][point.j])
	}

	return validValues
}

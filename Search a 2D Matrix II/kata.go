package kata

func searchMatrix(matrix [][]int, target int) bool {
outerLoop:
	for _, vector := range matrix {
		if len(vector) == 0 {
			continue
		}
		if vector[len(vector)/2] == target {
			return true
		}
		if vector[len(vector)/2] > target {
			for i := 0; i < len(vector)/2; i++ {
				if vector[i] > target {
					continue outerLoop
				}
				if vector[i] == target {
					return true
				}
			}
		} else if vector[len(vector)/2] < target {
			for i := len(vector) - 1; i > len(vector)/2; i-- {
				if vector[i] < target {
					continue outerLoop
				}
				if vector[i] == target {
					return true
				}
			}
		}
	}
	return false
}

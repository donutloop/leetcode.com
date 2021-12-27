package _4__Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 1 {
		return searchVector(matrix[0], target)
	}
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 {
			continue
		}
		if target == matrix[i][0] {
			return true
		}
		if matrix[i][0] < target && len(matrix)-1 != i {
			continue
		}
		row := i
		if matrix[i][0] > target && i != 0 {
			row = i - 1
		}
		return searchVector(matrix[row], target)
	}
	return false
}

func searchVector(v []int, target int) bool {

	if len(v) == 0 {
		return false
	}

	if len(v) == 1 {
		return v[0] == target
	}
	mid := len(v) / 2
	if v[mid] == target {
		return true
	}

	if v[mid] > target {
		return searchVector(v[:mid], target)
	}

	if v[mid] < target {
		return searchVector(v[mid:], target)
	}

	return false
}

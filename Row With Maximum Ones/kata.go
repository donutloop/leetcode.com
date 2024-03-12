package kata

func rowAndMaximumOnes(mat [][]int) []int {
	var globalMaxOnesPerRow int
	var rowIndex = 0
	for i, row := range mat {
		var localCountOnesPerRow int
		for _, col := range row {
			if col == 1 {
				localCountOnesPerRow++
			}
		}

		if localCountOnesPerRow > globalMaxOnesPerRow {
			globalMaxOnesPerRow = localCountOnesPerRow
			rowIndex = i
		}
	}
	return []int{rowIndex, globalMaxOnesPerRow}
}

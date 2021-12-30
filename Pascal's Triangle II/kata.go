package Pascal_s_Triangle_II

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	o := make([]int, 0)
	o = []int{1, 1}

	for i := 2; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[len(row)-1] = 1

		k := 1
		for j := 0; j < len(o)-1; j++ {
			n := o[j] + o[j+1]
			row[k] = n
			k++
		}
		o = row
	}

	return o
}

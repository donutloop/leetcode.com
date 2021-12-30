package Pascal_s_Triangle

func generate(numRows int) [][]int {

	if numRows == 1 {
		return [][]int{[]int{1}}
	}

	if numRows == 2 {
		return [][]int{[]int{1}, []int{1, 1}}
	}

	o := make([][]int, numRows)
	l := 0
	o[l] = []int{1}
	l++
	o[l] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[len(row)-1] = 1

		k := 1
		for j := 0; j < len(o[i-1])-1; j++ {
			n := o[i-1][j] + o[i-1][j+1]
			row[k] = n
			k++
		}
		l++
		o[l] = row
	}

	return o
}

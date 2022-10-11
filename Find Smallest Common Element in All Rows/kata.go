package Find_Smallest_Common_Element_in_All_Rows

func smallestCommonElement(mat [][]int) int {
	min := -1
	counters := make(map[int]int)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			c := counters[mat[i][j]]
			if c == i+1 {
				continue
			}

			c++
			if c == i+1 && len(mat)-1 == i {
				if min == -1 || mat[i][j] < min {
					min = mat[i][j]
				}
			}

			counters[mat[i][j]] = c
		}
	}
	return min
}

package kata

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			k := 0
			for j := i + 1; j < len(arr); j++ {
				l := arr[j]
				arr[j] = k
				k = l
			}
			i++
		}
	}
}

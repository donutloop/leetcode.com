package kata

func checkIfExist(arr []int) bool {

	doubles := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		doubles[arr[i]*2] = i
	}

	for i, n := range arr {
		if doubles[n] == i {
			continue
		}
		if _, ok := doubles[n]; ok {
			return true
		}
	}

	return false
}

package kata

func destCity(paths [][]string) string {
	start := make(map[string]bool)

	for _, v := range paths {
		start[v[0]] = true
	}

	var lastV string
	for _, v := range paths {
		if !start[v[1]] {
			lastV = v[1]
			break
		}
	}
	return lastV
}

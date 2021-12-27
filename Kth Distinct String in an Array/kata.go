package kata

func kthDistinct(arr []string, k int) string {

	distinct := make(map[string]int)
	for _, a := range arr {
		distinct[a]++
	}

	if len(distinct) < k {
		return ""
	}

	i := 1
	for _, a := range arr {
		if distinct[a] == 1 {
			if i == k {
				return a
			}
			i++
		}
	}

	return ""
}

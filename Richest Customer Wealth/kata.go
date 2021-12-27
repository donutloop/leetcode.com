package kata

func maximumWealth(accounts [][]int) int {
	var max int
	for _, account := range accounts {
		n := sum(account)
		if max < n {
			max = n
		}
	}
	return max
}

func sum(values []int) int {
	var sum int
	for i := 0; i < len(values)/2; i++ {
		sum += values[i] + values[len(values)-i-1]
	}
	if len(values)%2 == 1 {
		sum += values[len(values)/2]
	}

	return sum
}

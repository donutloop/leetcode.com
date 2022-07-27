package kata

func bestHand(ranks []int, suits []byte) string {

	c := make([]int, 13+4)
	for i, suit := range suits {
		c[ranks[i]-1]++
		c[suit-97+13]++
	}

	i := len(c) - 1
	for ; i >= 13; i-- {
		if c[i] == 5 {
			return "Flush"
		}
	}

	var max int
	for ; i >= 0; i-- {
		if c[i] > max {
			max = c[i]
		}
	}

	if max >= 3 {
		return "Three of a Kind"
	} else if max == 2 {
		return "Pair"
	}

	return "High Card"
}

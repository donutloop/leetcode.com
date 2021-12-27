package kata

func threeConsecutiveOdds(arr []int) bool {

	var c int
	for _, n := range arr {
		if n%2 == 1 {
			c++
			if c == 3 {
				return true
			}
		} else {
			c = 0
		}
	}

	return false
}

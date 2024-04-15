package kata

func scoreOfString(s string) int {
	var sum int
	for i := 1; i < len(s); i++ {
		sum += abs(int(s[i]) - int(s[i-1]))
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package kata

func numberOfSteps(num int) int {
	var steps int
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}

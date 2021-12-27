package kata

func totalMoney(n int) int {
	count := 0
	days := 0
	bouns := 0
	money := 1
	for i := 0; i < n; i++ {
		count += (money + bouns)
		days++
		money++
		if days == 7 {
			bouns++
			days = 0
			money = 1
		}
	}
	return count
}

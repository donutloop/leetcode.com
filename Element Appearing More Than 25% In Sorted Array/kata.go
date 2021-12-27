package kata

func findSpecialInteger(arr []int) int {
	counter := make(map[int]int)
	fragment := float64(len(arr)) / 100 * 25
	for _, n := range arr {
		counter[n]++
	}
	var match int
	max := -1
	for number, count := range counter {
		if float64(count) >= fragment && (max == -1 || max < count) {
			match = number
			max = count
		}
	}
	return match
}

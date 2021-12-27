package kata

func sumOfUnique(nums []int) int {
	set := make(map[int]int)

	for _, num := range nums {
		set[num]++
	}

	var sum int
	for num, count := range set {
		if count > 1 {
			continue
		}
		sum += num
	}

	return sum
}

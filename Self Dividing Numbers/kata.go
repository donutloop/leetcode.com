package kata

func selfDividingNumbers(left int, right int) []int {
	nums := make([]int, 0)
outerLoop:
	for i := left; i <= right; i++ {
		if i < 10 {
			nums = append(nums, i)
		} else {
			num := i
			for num > 0 {
				k := num % 10
				if k == 0 || i%k != 0 {
					continue outerLoop
				}
				num = num / 10
			}
			nums = append(nums, i)
		}
	}
	return nums
}

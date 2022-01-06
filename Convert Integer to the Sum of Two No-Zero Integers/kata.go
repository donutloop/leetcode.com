package Convert_Integer_to_the_Sum_of_Two_No_Zero_Integers

func getNoZeroIntegers(n int) []int {
	i := 1
	var k int
	for {
		k = n - i
		if !hasZeros(k) && !hasZeros(i) {
			return []int{i, k}
		}

		i++
	}
}

func hasZeros(num int) bool {
	for num > 0 {
		if num%10 == 0 {
			return true
		}
		num /= 10
	}
	return false
}

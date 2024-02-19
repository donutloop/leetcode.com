package main

func differenceOfSums(n int, m int) int {
	var sumDivisible, sumNotDivisible = 0, 0
	for num := 1; num <= n; num++ {
		if num%m == 0 {
			sumDivisible += num
		}
		if num%m != 0 {
			sumNotDivisible += num
		}
	}
	return sumNotDivisible - sumDivisible
}

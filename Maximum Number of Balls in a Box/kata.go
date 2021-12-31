package Maximum_Number_of_Balls_in_a_Box

func countBalls(lowLimit int, highLimit int) int {
	counter := make(map[int]int, highLimit-lowLimit+1)
	max := -1

	for i := lowLimit; i <= highLimit; i++ {

		n := i
		var sum int
		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		counter[sum]++
		if max == -1 || counter[sum] > max {
			max = counter[sum]
		}

	}

	return max
}

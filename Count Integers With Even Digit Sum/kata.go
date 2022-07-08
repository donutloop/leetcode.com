package Count_Integers_With_Even_Digit_Sum

var mem = map[int]int{}

func countEven(num int) int {

	c := 0
	i := 1
	for i <= num {
		e, ok := mem[i]
		if !ok {
			n := i
			var sum int
			for {
				k := n % 10
				sum += k
				n = n / 10
				if n == 0 {
					break
				}
			}
			mem[i] = sum
			if sum%2 == 0 {
				c++
			}
		} else {
			if e%2 == 0 {
				c++
			}
		}
		i++
	}

	return c
}

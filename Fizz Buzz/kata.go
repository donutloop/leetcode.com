package kata

import "strconv"

func fizzBuzz(n int) []string {
	values := make([]string, n)
	var j int
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			values[j] = "FizzBuzz"
		} else if i%3 == 0 {
			values[j] = "Fizz"
		} else if i%5 == 0 {
			values[j] = "Buzz"
		} else {
			values[j] = string(strconv.Itoa(i))
		}
		j++
	}
	return values
}

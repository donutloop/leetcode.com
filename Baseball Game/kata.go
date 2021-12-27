package kata

import "strconv"

func calPoints(ops []string) int {
	history := make([]int, 0)
	for _, op := range ops {

		if op == "C" {
			history = history[:len(history)-1]
			continue
		} else if op == "D" {
			n := history[len(history)-1] * 2
			history = append(history, n)
			continue
		} else if op == "+" {
			n := history[len(history)-1] + history[len(history)-2]
			history = append(history, n)
			continue
		}

		n, _ := strconv.Atoi(op)
		history = append(history, n)
	}

	var sum int
	for _, n := range history {
		sum = sum + n
	}

	return sum
}

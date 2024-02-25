package kata

func canSeePersonsCount(heights []int) []int {
	var stack []int
	result := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] > heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			result[i]++
		}
		if len(stack) >= 1 {
			result[i]++
		}
		stack = append(stack, i)
	}

	return result
}

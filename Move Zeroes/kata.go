package kata

func moveZeroes(nums []int) {
	stack := make([]int, 0)
	for _, num := range nums {
		if num == 0 {
			continue
		}
		stack = append(stack, num)
	}
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			nums[i] = 0
		} else {
			nums[i] = stack[0]
			stack = stack[1:]
		}
	}
}

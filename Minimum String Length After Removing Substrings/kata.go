package kata

func minLength(s string) int {
	if len(s) == 0 {
		return 0
	}

	var stack []byte
	var minimumLength = len(s)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && ((s[i] == 'B' && stack[len(stack)-1] == 'A') || (s[i] == 'D' && stack[len(stack)-1] == 'C')) {
			minimumLength -= 2
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return minimumLength
}

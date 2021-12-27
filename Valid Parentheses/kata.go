package kata

func isValid(s string) bool {
	stack := make([]rune, 0)

	if len(s)%2 == 1 {
		return false
	}

	var openChar rune
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
			continue
		}

		openChar = rune(1)
		if len(stack)-1 > -1 {
			openChar = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if char == '}' && openChar != '{' || char == ')' && openChar != '(' || char == ']' && openChar != '[' {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

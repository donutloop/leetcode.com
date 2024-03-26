package kata

func makeGood(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if len(stack) > 0 && stack[len(stack)-1] >= 'A' && stack[len(stack)-1] <= 'Z' {
				if (stack[len(stack)-1] + 32) == s[i] {
					stack = stack[:len(stack)-1]
					continue
				}
				stack = append(stack, s[i])
				continue
			}
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == (s[i]+32) {
				stack = stack[:len(stack)-1]
				continue
			}
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

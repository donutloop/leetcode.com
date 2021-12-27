package kata

const o byte = 'o'

func interpret(command string) string {
	var s []byte
	for i := 0; i < len(command); i++ {
		if isOpen(command[i]) && i+1 < len(command) && isClosed(command[i+1]) {
			s = append(s, o)
			i++
		} else if !isOpen(command[i]) && !isClosed(command[i]) {
			s = append(s, command[i])
		} else if isClosed(command[i]) {
			continue
		}
	}
	return string(s)
}

func isClosed(c byte) bool {
	if c == ')' {
		return true
	}
	return false
}

func isOpen(c byte) bool {
	if c == '(' {
		return true
	}
	return false
}

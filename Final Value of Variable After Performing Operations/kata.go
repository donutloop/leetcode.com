package kata

func finalValueAfterOperations(operations []string) int {

	x := 0

	for _, operation := range operations {

		if len(operation) != 3 {
			panic("operation invalid")
		}

		if operation[0] == '-' {
			if operation[1] == '-' {
				if operation[2] == 'X' {
					x = x - 1
					continue
				} else {
					panic("operation invalid")
				}
			} else {
				panic("operation invalid")
			}
		}

		if operation[0] == '+' {
			if operation[1] == '+' {
				if operation[2] == 'X' {
					x = x + 1
					continue
				} else {
					panic("operation invalid")
				}
			} else {
				panic("operation invalid")
			}
		}

		if operation[0] == 'X' {
			if operation[1] == '+' {
				if operation[2] == '+' {
					x = x + 1
					continue
				} else {
					panic("operation invalid")
				}
			} else if operation[1] == '-' {
				if operation[2] == '-' {
					x = x - 1
					continue
				} else {
					panic("operation invalid")
				}
			} else {
				panic("operation invalid")
			}
		} else {
			panic("operation invalid")
		}

	}

	return x
}

package kata

const (
	RightShiftMode int = 1
	LeftShiftMode  int = 0
)

func stringShift(s string, shift [][]int) string {
	b := []byte(s)
	for _, operation := range shift {
		amount := operation[1]
		shiftMode := operation[0]
		if shiftMode == RightShiftMode {
			b = RightShift(b, amount)
		} else if shiftMode == LeftShiftMode {
			b = LeftShift(b, amount)
		}
	}
	return string(b)
}

func LeftShift(b []byte, amount int) []byte {
	if amount == 0 {
		return b
	}

	if len(b) <= 2 {
		return b
	}

	tmp := b[0]
	for i := 0; i < len(b)-1; i++ {
		b[i] = b[i+1]
	}
	b[len(b)-1] = tmp

	return LeftShift(b, amount-1)
}

func RightShift(b []byte, amount int) []byte {
	if amount == 0 {
		return b
	}

	if len(b) <= 2 {
		return b
	}

	tmp := b[len(b)-1]
	for i := len(b) - 1; i >= 1; i-- {
		b[i] = b[i-1]
	}
	b[0] = tmp

	return RightShift(b, amount-1)
}

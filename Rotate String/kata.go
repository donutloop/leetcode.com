package kata

func rotateString(A string, B string) bool {
	if A == "" && B == "" {
		return true
	}
	if A == "" {
		return false
	}

	oldA := A
	first := true
	for {
		if first != true && oldA == A {
			return false
		}
		first = false

		a := A[0]
		A = A[1:len(A)] + string(a)

		if A == B {
			return true
		}
	}

	return false
}

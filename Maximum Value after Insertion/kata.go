package Maximum_Value_after_Insertion

func maxValue(n string, x int) string {

	b := make([]byte, len(n)+1)

	y := byte(x + 48)

	var i int
	if n[0] == '-' {
		i = 1
		b[0] = '-'
	}

	ok := true

	if i == 0 {
		k := i
		for j := i; j < len(n); j++ {
			if ok && y > n[j] {
				b[k] = y
				k++
				b[k] = n[j]
				ok = false
			} else {
				b[k] = n[j]
			}
			k++
		}
	} else {
		k := i
		for j := i; j < len(n); j++ {
			if ok && y < n[j] {

				b[k] = y
				k++
				b[k] = n[j]
				ok = false
			} else {
				b[k] = n[j]
			}
			k++
		}

	}

	if ok {
		b[len(b)-1] = y
	}

	return string(b)
}

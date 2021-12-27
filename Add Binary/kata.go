package kata

func addBinary(a string, b string) string {

	if len(a) == 0 && len(b) == 0 {
		return "0"
	}

	bb := []byte(b)
	aa := []byte(a)

	if len(aa) > len(bb) {
		for len(bb) < len(aa) {
			bb = append([]byte{'0'}, bb...)
		}
	} else {
		for len(aa) < len(bb) {
			aa = append([]byte{'0'}, aa...)
		}
	}

	carry := byte(0)
	for i := len(bb) - 1; i >= 0; i-- {
		if (bb[i] == '0' && aa[i] == '0') && carry == byte(1) {
			bb[i] = '1'
			carry = byte(0)
		} else if (bb[i] == '1' && aa[i] == '1') && carry == byte(1) {
			bb[i] = '1'
		} else if (bb[i] == '1' || aa[i] == '1') && carry == byte(1) {
			bb[i] = '0'
		} else if bb[i] == '1' && aa[i] == '1' {
			bb[i] = '0'
			carry = byte(1)
		} else if (bb[i] == '0' && aa[i] == '1') || (bb[i] == '1' && aa[i] == '0') {
			bb[i] = '1'
			carry = byte(0)
		} else if bb[i] == '0' && aa[i] == '0' {
			bb[i] = '0'
			carry = byte(0)
		}
	}

	if carry == byte(1) {
		bb = append([]byte{'1'}, bb...)
	}

	return string(bb)
}

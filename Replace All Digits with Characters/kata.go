package kata

func replaceDigits(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i = i + 2 {
		if i+1 == len(b) {
			break
		}
		b[i+1] = shift(b[i], b[i+1])
	}
	return string(b)
}

func shift(c byte, by byte) byte {
	return c + by - 48
}

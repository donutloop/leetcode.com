package kata

func toLowerCase(str string) string {
	if len(str) == 0 {
		return str
	}
	b := []byte(str)
	var i int
	for _, c := range b {
		if c >= 65 && c <= 90 {
			b[i] = b[i] + 32
		}
		i++
	}
	return string(b)
}

package kata

func generateTheString(n int) string {
	s := make([]byte, n)

	var i int
	if n%2 == 1 {
		i = 0
	} else {
		s[0] = byte('a')
		i = 1
	}

	b := byte('b')
	for ; i < n; i++ {
		s[i] = b
	}

	return string(s)
}

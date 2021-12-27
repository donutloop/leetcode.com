package kata

func reverseWords(s string) string {

	b := []byte(s)
	var k int
	for i := 0; i < len(b); i++ {

		if b[i] == 32 || len(b) == (i+1) {

			if len(b) == (i + 1) {
				i = i + 1
			}

			j := i - 1
			n := (i - k) % 2
			var m int
			if n == 0 {
				m = (i - 1 - k) / 2
			} else {
				m = (i - 0 - k) / 2
			}

			for l := 0; l <= m; l++ {
				b[k], b[j] = b[j], b[k]
				j--
				k++
			}
			k = i + 1
		}
	}

	return string(b)

}

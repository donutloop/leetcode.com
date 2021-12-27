package kata

import "strconv"

func thousandSeparator(n int) string {
	s := strconv.Itoa(n)
	k := make([]byte, 0)
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if count == 3 {
			k = append(k, byte(46))
			count = 0
		}
		count++
		k = append(k, s[i])
	}
	l := make([]byte, len(k))
	j := 0
	for i := len(k) - 1; i >= 0; i-- {
		l[j] = k[i]
		j++
	}
	return string(l)
}

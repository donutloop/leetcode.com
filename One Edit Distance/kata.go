package One_Edit_Distance

import "math"

func isOneEditDistance(s string, t string) bool {
	if s == t {
		return false
	}

	if math.Abs(float64(len(s))-float64(len(t))) > 1 {
		return false
	}

	var c int
	if len(s) == len(t) {
		for i := 0; i < len(t); i++ {
			if s[i] != t[i] {
				c++
			}
		}
		return c == 1
	}

	if len(s) > len(t) {
		if len(t) == 0 {
			return true
		}
		if s[:len(s)-1] == t {
			return true
		}
		j := 0
		for i := 0; i < len(s); i++ {
			if s[i] != t[j] {
				c++
				continue
			}
			j++
		}
		return c == 1 || c == 0
	}

	if string(t[0])+s == t {
		return true
	}
	if s+string(t[len(t)-1]) == t {
		return true
	}

	j := 0
	for i := 0; i < len(t); i++ {
		if s[j] != t[i] {
			c++
			continue
		}
		j++
	}
	return c == 1 || c == 0
}

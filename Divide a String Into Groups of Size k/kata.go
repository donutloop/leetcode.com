package Divide_a_String_Into_Groups_of_Size_k

func divideString(s string, k int, fill byte) []string {

	res := make([]string, 0)
	if k > len(s) {
		m := len(s)
		for l := 0; l < k-m; l++ {
			s += string(fill)
		}
		res = append(res, s)
		return res
	}

	j := 0
	i := k
	for i <= len(s) {
		res = append(res, s[j:i])
		j = i
		i = i + k
		if i > len(s) && len(s[j:]) > 0 {
			b := make([]byte, 0)
			b = append(b, s[j:]...)
			m := len(b)
			for l := 0; l < k-m; l++ {
				b = append(b, fill)
			}
			res = append(res, string(b))
			break
		}
	}

	return res
}

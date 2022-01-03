package Delete_Characters_to_Make_Fancy_String

func makeFancyString(s string) string {
	c := 1
	char := s[0]
	b := make([]byte, 0)
	b = append(b, s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == char {
			c++
			if c > 2 {
				continue
			}
		} else {
			c = 1
			char = s[i]
		}
		b = append(b, s[i])
	}
	return string(b)
}

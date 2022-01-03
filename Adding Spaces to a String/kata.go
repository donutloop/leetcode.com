package Adding_Spaces_to_a_String

func addSpaces(s string, spaces []int) string {
	b := make([]byte, len(s)+len(spaces))
	var j int
	var k int
	for i := 0; i < len(s)+len(spaces); i++ {
		if k < len(spaces) && spaces[k] == j {
			b[i] = ' '
			k++
		} else {
			b[i] = s[j]
			j++
		}
	}
	return string(b)
}

package Check_if_Binary_String_Has_at_Most_One_Segment_of_Ones

func checkOnesSegment(s string) bool {

	if len(s) == 1 && s[0] == '1' {
		return true
	}

	c := 0
	oneContiguousSegment := false

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			if oneContiguousSegment {
				return false
			}

			c++
		} else {
			if c >= 1 {
				oneContiguousSegment = true
			}

			c = 0
		}
	}

	if c > 0 {
		return true
	}

	return oneContiguousSegment
}

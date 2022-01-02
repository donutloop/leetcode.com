package Check_if_All_A_s_Appears_Before_All_B_s

func checkString(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == 'b' && s[i] == 'a' {
			return false
		}
	}
	return true
}

package A_Number_After_a_Double_Reversal

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	return num%10 > 0
}

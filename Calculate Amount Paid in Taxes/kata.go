package Calculate_Amount_Paid_in_Taxes

func calculateTax(brackets [][]int, income int) float64 {

	if income == 0 {
		return 0
	}

	k := 0
	amount := 0.0
	for _, bracket := range brackets {

		if income >= bracket[0] {
			if k == 0 {
				k += bracket[0]
				amount += float64(bracket[0]) * float64(bracket[1]) / 100
			} else {
				amount += (float64(bracket[0]) - float64(k)) * float64(bracket[1]) / 100
				k += bracket[0] - k
			}
		} else if income < bracket[0] {
			amount += float64(income-k) * float64(bracket[1]) / 100
			break
		}
	}

	return amount
}

package Largest_3_Same_Digit_Number_in_String

func largestGoodInteger(num string) string {

	k := -1
	window := ""
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i-1] == num[i-2] && k < int(num[i]-48) {
			window = num[i-2 : i+1]
			k = int(num[i] - 48)
		}
	}

	return window
}

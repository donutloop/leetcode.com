package _812__Determine_Color_of_a_Chessboard_Square

func squareIsWhite(coordinates string) bool {
	a := int(coordinates[0]) - 96
	b := int(coordinates[1]) - 48
	if a%2 == 1 {
		if b%2 == 0 {
			return true
		} else {
			return false
		}
	} else {
		if b%2 == 0 {
			return false
		} else {
			return true
		}
	}
}

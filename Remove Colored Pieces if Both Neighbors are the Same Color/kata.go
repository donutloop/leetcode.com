package Remove_Colored_Pieces_if_Both_Neighbors_are_the_Same_Color

func winnerOfGame(colors string) bool {

	if len(colors) < 3 {
		return false
	}

	cb := []byte(colors)

	a := 0
	b := 0

	for i := 2; i < len(cb); i++ {
		if cb[i-2] == 'A' && cb[i-1] == 'A' && cb[i] == 'A' {
			a++
			cb[i-2] = 'N'
		} else if cb[i-2] == 'B' && cb[i-1] == 'B' && cb[i] == 'B' {
			b++
			cb[i-2] = 'N'
		}
	}

	if a == b {
		return false
	}

	return a > b
}

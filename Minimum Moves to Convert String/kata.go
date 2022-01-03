package Minimum_Moves_to_Convert_String

func minimumMoves(s string) int {
	var c int
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			c++
			i = i + 2
		}
	}
	return c
}

package Remove_Vowels_from_a_String

func removeVowels(s string) string {
	b := make([]rune, 0, len(s)/3)
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			continue
		}
		b = append(b, char)
	}
	return string(b)
}

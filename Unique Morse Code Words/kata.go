package kata

var encoder []string = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {

	uniqueTransformations := map[string]bool{}
	for _, word := range words {
		var encodedWord string
		for _, char := range word {
			encodedWord += encoder[int(char-97)]
		}
		uniqueTransformations[encodedWord] = true
	}

	return len(uniqueTransformations)
}

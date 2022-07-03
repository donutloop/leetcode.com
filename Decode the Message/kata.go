package Decode_the_Message

func decodeMessage(key string, message string) string {

	decoderTable := map[rune]rune{}

	j := rune(97)

	for _, char := range key {
		if char == 32 {
			continue
		}

		_, ok := decoderTable[char]
		if !ok {
			decoderTable[char] = j
			j++
		}
	}

	decodedMessage := make([]rune, len(message))
	for i, char := range message {
		if char == 32 {
			decodedMessage[i] = char
		} else {
			decodedMessage[i] = decoderTable[char]
		}
	}

	return string(decodedMessage)
}

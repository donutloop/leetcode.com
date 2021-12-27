package kata

func defangIPaddr(address string) string {
	defangAddress := make([]rune, 0, len(address)+6)
	for _, char := range address {
		if char == '.' {
			defangAddress = append(defangAddress, '[', '.', ']')
			continue
		}
		defangAddress = append(defangAddress, char)
	}
	return string(defangAddress)
}

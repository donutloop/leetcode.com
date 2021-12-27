package kata

func backspaceCompare(S string, T string) bool {
	return mod(S) == mod(T)
}

func mod(S string) string {
	modifiedS := make([]byte, 0)
	var j int
	for i := 0; i < len(S); i++ {
		if S[i] == '#' {
			if j-1 >= 0 {
				modifiedS = modifiedS[:j-1]
				j = j - 1
			}
			continue
		}
		j++
		modifiedS = append(modifiedS, S[i])
	}

	return string(modifiedS)
}

package kata

func letterCasePermutation(S string) []string {
	permutations := make([]string, 0)
	permutations = append(permutations, string(permutation(&permutations, []byte(S), 0)))
	return permutations
}

func permutation(permutations *[]string, b []byte, index int) []byte {
	if len(b) <= index {
		return b
	}
	for i := index; i < len(b); i++ {
		if b[i] >= 97 && b[i] <= 122 {
			b[i] = b[i] - 32
			*permutations = append(*permutations, string(permutation(permutations, b, i+1)))
			b[i] = b[i] + 32
		} else if b[i] >= 65 && b[i] <= 90 {
			b[i] = b[i] + 32
			*permutations = append(*permutations, string(permutation(permutations, b, i+1)))
			b[i] = b[i] - 32
		}
	}
	return b
}

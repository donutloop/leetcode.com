package kata

func uncommonFromSentences(A string, B string) []string {
	type WordCounter struct {
		A int
		B int
	}
	raw := make([]byte, 0)
	set := make(map[string]WordCounter)
	for i := 0; i < len(A); i++ {
		if A[i] == ' ' || len(A)-1 == i {
			if len(A)-1 == i {
				raw = append(raw, A[i])
			}

			word := string(raw)
			v, ok := set[word]
			if ok {
				v.A = v.A + 1
				set[word] = v
			} else {
				set[word] = WordCounter{A: 1}
			}
			raw = make([]byte, 0)
			continue
		}
		raw = append(raw, A[i])
	}
	raw = make([]byte, 0)
	for i := 0; i < len(B); i++ {
		if B[i] == ' ' || len(B)-1 == i {
			if len(B)-1 == i {
				raw = append(raw, B[i])
			}

			word := string(raw)
			v, ok := set[word]
			if ok {
				v.B = v.B + 1
				set[word] = v
			} else {
				set[word] = WordCounter{B: 1}
			}
			raw = make([]byte, 0)
			continue
		}
		raw = append(raw, B[i])
	}
	words := make([]string, 0)
	for w, c := range set {
		if c.A == 1 && c.B == 0 || c.A == 0 && c.B == 1 {
			words = append(words, w)
		}
	}
	return words
}

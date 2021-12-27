package kata

import "sort"

const whitespace = ' '

func sortSentence(s string) string {

	type word struct {
		idx   int
		chars []byte
	}

	words := make([]word, 0)
	w := word{idx: -1, chars: make([]byte, 0)}
	for i := 0; i < len(s); i++ {
		if (i+1 < len(s) && whitespace == s[i+1]) || len(s)-1 == i {
			idx := s[i] - 48
			i++
			w.idx = int(idx)
			words = append(words, w)
			w = word{idx: -1, chars: make([]byte, 0)}
			continue
		}
		w.chars = append(w.chars, s[i])
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].idx < words[j].idx
	})

	sortedS := make([]byte, 0)
	for _, word := range words {
		sortedS = append(sortedS, word.chars...)
		sortedS = append(sortedS, ' ')
	}

	return string(sortedS[:len(sortedS)-1])
}

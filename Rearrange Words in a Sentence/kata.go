package kata

import "sort"

func arrangeWords(text string) string {
	if len(text) == 0 {
		return text
	}

	type Word struct {
		value []byte
		idx   int
	}

	word := make([]byte, 0)
	wordEmpty := make([]byte, 0)
	wordList := make([]Word, 0)
	var wordCounter int
	for i := 0; i < len(text); i++ {
		if text[i] != ' ' {
			word = append(word, text[i])
		}

		if text[i] == ' ' || len(text)-1 == i {
			wordList = append(wordList, Word{value: word, idx: wordCounter})
			word = wordEmpty
			wordCounter++
		}
	}

	sort.Slice(wordList, func(i, j int) bool {
		if len(wordList[i].value) == len(wordList[j].value) && wordList[i].idx < wordList[j].idx {
			return true
		} else if len(wordList[i].value) < len(wordList[j].value) {
			return true
		}
		return false
	})

	if len(wordList) == 0 {
		return text
	}

	if wordList[0].value[0] >= 97 && wordList[0].value[0] <= 122 {
		wordList[0].value[0] = wordList[0].value[0] - 32
	}

	text = ""
	text += string(wordList[0].value) + " "

	for _, word := range wordList[1:] {
		if word.value[0] >= 65 && word.value[0] <= 90 {
			word.value[0] = word.value[0] + 32
		}

		text += string(word.value) + " "
	}
	return text[:len(text)-1]
}

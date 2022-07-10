package Implement_Magic_Dictionary

type MagicDictionary struct {
	words map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		words: nil,
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	words := make(map[int][]string, 0)
	for _, word := range dictionary {
		_, ok := words[len(word)]
		if !ok {
			words[len(word)] = []string{}
		}

		words[len(word)] = append(words[len(word)], word)
	}

	this.words = words
}

func (this *MagicDictionary) Search(searchWord string) bool {
	ws, ok := this.words[len(searchWord)]
	if !ok {
		return false
	}
outer:
	for _, word := range ws {
		if word == searchWord {
			continue
		}

		if len(word) == 1 {
			return true
		}

		c := 0
		for i := 0; i < len(searchWord); i++ {
			if searchWord[i] != word[i] {
				c++
			}
			if c > 1 {
				continue outer
			}
		}
		return true
	}

	return false
}

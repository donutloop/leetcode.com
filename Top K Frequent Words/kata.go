package kata

import "sort"

type Prio struct {
	word string
	freq int
}

func topKFrequent(words []string, k int) []string {
	if len(words) == 0 {
		return nil
	}
	counter := make(map[string]int)
	for i := range words {
		counter[words[i]]++
	}

	prios := make([]Prio, len(counter))
	var i int
	for w, c := range counter {
		prio := Prio{
			word: w,
			freq: c,
		}
		prios[i] = prio
		i++
	}
	sort.Slice(prios, func(i, j int) bool {
		if prios[i].freq == prios[j].freq {
			return prios[i].word < prios[j].word
		}
		return prios[i].freq > prios[j].freq
	})

	prios = prios[:k]
	matchedWords := make([]string, len(prios))
	for i := range prios {
		matchedWords[i] = prios[i].word
	}
	return matchedWords
}

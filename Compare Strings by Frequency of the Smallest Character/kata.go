package kata

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	nums := make([]int, len(queries))
	memory := make(map[string]int)
	for i, query := range queries {
		var count int
		queryFrequency := f([]byte(query))
		for _, word := range words {
			wordFrequency, ok := memory[word]
			if !ok {
				wordFrequency = f([]byte(word))
				memory[word] = wordFrequency
			}

			if wordFrequency > queryFrequency {
				count++
			}
		}
		nums[i] = count
	}
	return nums
}

func f(s []byte) int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	lastChar := s[0]
	c := 1
	for i := 1; i < len(s); i++ {
		if lastChar != s[i] {
			return c
		}
		c++
		lastChar = s[i]
	}
	return c
}

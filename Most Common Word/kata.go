package kata

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	if len(paragraph) == 0 {
		return ""
	}
	words := make(map[string]int)
	var word string
	for i := 0; i < len(paragraph); i++ {
		if (paragraph[i] >= 65 && paragraph[i] <= 90) || (paragraph[i] >= 97 && paragraph[i] <= 122) {
			word += string(paragraph[i])
			if len(paragraph)-1 != i {
				continue
			}
		}
		if word == "" {
			continue
		}
		word = strings.ToLower(word)
		if !hasBanned(banned, word) {
			words[word] = words[word] + 1
		}
		word = ""
	}
	if len(words) == 0 {
		return strings.ToLower(word)
	}
	max := -1
	for w, c := range words {
		fmt.Println(w, c)
		if max == -1 || c > max {
			word = w
			max = c
		}
	}
	return word
}

func hasBanned(banned []string, word string) bool {
	for i := 0; i < len(banned); i++ {
		if strings.ToLower(banned[i]) == word {
			return true
		}
	}
	return false
}

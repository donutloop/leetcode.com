package Shortest_Word_Distance

import "math"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	k, j := -1, -1
	globalDiff := -1
	for i, word := range wordsDict {
		if word == word1 {
			k = i
			if k > -1 && j > -1 {
				localDiff := int(math.Abs(float64(k) - float64(j)))
				if globalDiff == -1 || localDiff < globalDiff {
					globalDiff = localDiff
				}
			}
		} else if word == word2 {
			j = i
			if k > -1 && j > -1 {
				localDiff := int(math.Abs(float64(k) - float64(j)))
				if globalDiff == -1 || localDiff < globalDiff {
					globalDiff = localDiff
				}
			}
		}
	}
	return globalDiff
}

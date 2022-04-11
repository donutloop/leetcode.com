package Shortest_Word_Distance_II

import "math"

type WordDistance struct {
	wordsDict []string
}

func Constructor(wordsDict []string) WordDistance {
	return WordDistance{
		wordsDict: wordsDict,
	}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	return shortestDistance(this.wordsDict, word1, word2)
}

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

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */

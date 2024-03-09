package kata

import "sort"

type player struct {
	losses int
	wins   int
}

func findWinners(matches [][]int) [][]int {
	allPlayers := make(map[int]player)
	for _, match := range matches {
		player := allPlayers[match[1]]
		player.losses++
		allPlayers[match[1]] = player

		player = allPlayers[match[0]]
		player.wins++
		allPlayers[match[0]] = player
	}

	var exactlyOneMatch []int
	var notLostAnyMatches []int

	for player, stats := range allPlayers {
		if stats.losses == 0 {
			notLostAnyMatches = append(notLostAnyMatches, player)
		} else if stats.losses == 1 {
			exactlyOneMatch = append(exactlyOneMatch, player)
		}
	}

	sort.Ints(notLostAnyMatches)
	sort.Ints(exactlyOneMatch)

	return [][]int{notLostAnyMatches, exactlyOneMatch}
}

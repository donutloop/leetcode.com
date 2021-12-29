package Rings_and_Rods

import "sort"

func countPoints(rings string) int {

	pairs := make([]Pair, len(rings)/2)
	var j int
	for i := 0; i < len(rings); i = i + 2 {
		pairs[j] = Pair{
			Key:   rings[i+1],
			Value: rings[i],
		}
		j++
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Key == pairs[j].Key {
			return pairs[i].Value < pairs[j].Value
		}
		return pairs[i].Key < pairs[j].Key
	})

	currentRode := pairs[0].Key
	currentColor := pairs[0].Value
	counter := 1
	all := 0
	for i := 1; i < len(pairs); i++ {
		if currentRode != pairs[i].Key || (len(pairs)-1 == i) {

			if len(pairs)-1 == i && currentColor != pairs[i].Value && pairs[i].Key == currentRode {
				counter++
			}

			if counter == 3 {
				all++
			}

			currentRode = pairs[i].Key
			currentColor = pairs[i].Value
			counter = 1
			continue
		}

		if currentColor != pairs[i].Value {
			counter++
			currentColor = pairs[i].Value
		}

	}

	return all
}

type Pair struct {
	Key   byte
	Value byte
}

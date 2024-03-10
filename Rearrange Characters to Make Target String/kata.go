package kata

func rearrangeCharacters(s string, target string) int {
	targetCharStats := [26]int{}
	for i := range target {
		targetCharStats[normalise(int(target[i]))]++
	}

	sCharStats := [26]int{}
	for i := range s {
		sCharStats[normalise(int(s[i]))]++
	}

	targetCount := -1
	for i := 0; i < 26; i++ {
		if (sCharStats[i] == 0 && targetCharStats[i] == 0) ||
			(sCharStats[i] > 0 && targetCharStats[i] == 0) {
			continue
		}
		if targetCharStats[i] > 0 && sCharStats[i] == 0 {
			return 0
		}

		if sCharStats[i] >= targetCharStats[i] {
			possibleTargetCount := float64(sCharStats[i]) / float64(targetCharStats[i])

			if targetCount == -1 || int(possibleTargetCount) < targetCount {
				targetCount = int(possibleTargetCount)
			}
		} else {
			return 0
		}
	}

	return targetCount
}

func normalise(a int) int {
	return a - 97
}

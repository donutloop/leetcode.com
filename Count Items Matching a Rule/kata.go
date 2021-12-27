package kata

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	idx := 0
	if ruleKey == "type" {
		idx = 0
	} else if ruleKey == "color" {
		idx = 1
	} else if ruleKey == "name" {
		idx = 2
	}

	var count int
	for _, item := range items {
		if item[idx] == ruleValue {
			count++
		}
	}

	return count
}

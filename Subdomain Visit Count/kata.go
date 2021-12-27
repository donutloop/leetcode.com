package kata

import "strconv"

func subdomainVisits(cpdomains []string) []string {
	if len(cpdomains) == 0 {
		return make([]string, 0)
	}

	stats := make(map[string]int)
	for _, u := range cpdomains {
		levelBIndex := -1
		levelAIndex := -1
		countDot := 0
		i := 0
		for ; i < len(u); i++ {
			if u[i] == ' ' {
				break
			}
		}
		i++
		domainStartIndex := i
		c, _ := strconv.Atoi(string(u[:i-1]))

		markedA := false

		for ; i < len(u); i++ {

			if u[i] == '.' {
				countDot++
			}

			if countDot == 1 && !markedA {
				levelAIndex = i + 1
				markedA = true
			}

			if countDot == 2 {
				levelBIndex = i + 1
				break
			}

		}

		stats[string(u[domainStartIndex:])] += c
		if levelAIndex != -1 {
			stats[u[levelAIndex:]] += c
		}
		if levelBIndex != -1 {
			stats[u[levelBIndex:]] += c
		}
	}

	res := make([]string, len(stats))

	i := 0
	for id, stat := range stats {
		res[i] = strconv.Itoa(stat) + " " + id
		i++
	}

	return res
}

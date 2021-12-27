package kata

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	toDelete := make(map[int]bool)
	for {
		ok := merge(accounts, toDelete)
		if !ok {
			break
		}
	}

	previousLen := len(accounts)
	for k := 0; k < previousLen; k++ {
		if !toDelete[k] {
			accounts[k] = emailUnion(accounts[k], 1)
			sort.Slice(accounts[k], func(i, j int) bool {
				return accounts[k][i] < accounts[k][j]
			})
			accounts = append(accounts, accounts[k])
		}
	}
	return accounts[previousLen:]
}

func merge(accounts [][]string, toDelete map[int]bool) bool {
	emails := make(map[string]int)
	for i, account := range accounts {
		if toDelete[i] {
			continue
		}

		for _, email := range account[1:] {
			v, ok := emails[email]
			if ok && emails[email] != i {
				accounts[v] = append(accounts[v], accounts[i][1:]...)
				toDelete[i] = true
				return true
			}

			emails[email] = i
		}
	}
	return false
}

func emailUnion(a []string, startIdx int) []string {
	set := make(map[string]bool, len(a)/2)
	for i := startIdx; i < len(a); i++ {
		set[a[i]] = true
	}
	a = a[:1]
	for elem := range set {
		a = append(a, elem)
	}
	return a
}

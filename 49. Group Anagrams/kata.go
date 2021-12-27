package _9__Group_Anagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	set := make(map[string][]string)
	lists := make([][]string, 0)
	for _, s := range strs {
		sb := []byte(s)
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})
		list, ok := set[string(sb)]
		if ok {
			list = append(list, s)
			set[string(sb)] = list
		} else {
			l := []string{s}
			set[string(sb)] = l
		}
	}
	for _, list := range set {
		lists = append(lists, list)
	}

	return lists
}

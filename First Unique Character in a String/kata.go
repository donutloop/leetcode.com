package kata

func firstUniqChar(s string) int {
	type counter struct {
		i int
		c int
	}

	set := make(map[rune]*counter)
	for i, c := range s {
		v, ok := set[c]
		if !ok {
			set[c] = &counter{c: 1, i: i}
			continue
		}
		v.c = v.c + 1
		set[c] = v
	}

	index := -1
	for _, v := range set {
		if (index == -1 && v.c == 1) || (v.c == 1 && index > v.i) {
			index = v.i
		}
	}
	return index
}

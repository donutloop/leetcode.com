package Greatest_English_Letter_in_Upper_and_Lower_Case

import "sort"

func greatestLetter(s string) string {

	b := []byte(s)

	sort.Slice(b, func(i, j int) bool {
		bj, bi := b[i], b[j]
		if bi >= 97 && bi <= 122 {
			bi = bi - 32
		}
		if bj >= 97 && bj <= 122 {
			bj = bj - 32
		}
		return bi > bj
	})

	i := len(b) - 2
	for j := len(b) - 1; j > 0; j-- {

		bj, bi := b[i], b[j]

		if bj == bi {
			continue
		}

		if bi >= 97 && bi <= 122 {
			if bi-32 == bj {
				return string(bj)
			}
		}

		if bj >= 97 && bj <= 122 {
			if bi == bj-32 {
				return string(bi)
			}
		}
		i--
	}

	return ""

}

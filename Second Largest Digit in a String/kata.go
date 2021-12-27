package kata

import "sort"

func secondHighest(s string) int {

	b := []byte(s)

	sort.Slice(b, func(i, j int) bool {
		bi := int(b[i])
		bj := int(b[j])

		if !(bi >= 48 && bi <= 57) {
			bi = -9999
		}

		if !(bj >= 48 && bj <= 57) {
			bj = -9999
		}

		return bi < bj
	})

	max := b[len(b)-1]
	for i := len(b) - 2; i >= 0; i-- {
		if max != b[i] {
			if b[i] >= 48 && b[i] <= 57 {
				return int(b[i] - 48)
			}
			break
		}
	}

	return -1
}

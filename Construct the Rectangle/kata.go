package Construct_the_Rectangle

import "math"

func constructRectangle(area int) []int {
	if area == 1 {
		return []int{1, 1}
	}

	nl, w, l := area, 0, 0
	diff := math.Inf(1)
	for {
		if area%nl == 0 {
			nw := area / nl
			if nw > nl {
				break
			}

			if diff > float64(nl-nw) {

				w = nw
				l = nl
				diff = float64(nl - nw)
			}

		}
		nl--
	}
	return []int{l, w}
}

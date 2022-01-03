package Two_Furthest_Houses_With_Different_Colors

func maxDistance(colors []int) int {
	a := colors[0]
	b := colors[len(colors)-1]
	left := -1
	right := -1
	var i int
	for j := len(colors) - 1; j >= 1; j-- {
		if left == -1 && a != colors[j] {
			left = j
		}
		if right == -1 && b != colors[i] {
			right = j
		}
		i++
	}

	if left > right {
		return left
	}

	return right
}

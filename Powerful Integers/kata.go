package kata

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	i := 0
	j := 0

	set := make(map[int]bool)
	var hit bool
outerLoop:
	for {
		xi := int(math.Pow(float64(x), float64(i)))
		if xi >= bound || (hit && x == 1) {
			break
		}
		i++
		for {
			hit = true
			yj := int(math.Pow(float64(y), float64(j)))
			z := xi + yj
			if z > bound {
				j = 0
				continue outerLoop
			}

			set[z] = true
			if y == 1 {
				j = 0
				continue outerLoop
			}

			j++
		}
	}

	nums := make([]int, len(set))
	i = 0
	for z := range set {
		nums[i] = z
		i++
	}

	return nums
}

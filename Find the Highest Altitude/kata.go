package Find_the_Highest_Altitude

func largestAltitude(gain []int) int {
	var current int
	highestAltitude := current
	for _, n := range gain {
		current += n
		if current > highestAltitude {
			highestAltitude = current
		}
	}
	return highestAltitude
}

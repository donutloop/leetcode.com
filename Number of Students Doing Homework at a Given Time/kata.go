package kata

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var matched int
	for i := 0; i < len(startTime); i++ {
		if startTime[i] > queryTime {
			continue
		}
		if endTime[i] < queryTime {
			continue
		}
		matched++
	}
	return matched
}

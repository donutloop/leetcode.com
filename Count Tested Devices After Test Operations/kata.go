package kata

func countTestedDevices(batteryPercentages []int) int {
	var counter int
	for i := 0; i < len(batteryPercentages); i++ {
		if (batteryPercentages[i] - counter) == 0 {
			continue
		} else if (batteryPercentages[i] - counter) > 0 {
			counter++
		}
	}
	return counter
}

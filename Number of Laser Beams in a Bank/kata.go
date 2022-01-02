package Number_of_Laser_Beams_in_a_Bank

func numberOfBeams(bank []string) int {
	var sum int
	preCurrentDevices := -1
	for i := 0; i < len(bank); i++ {
		var currentDevices int
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				currentDevices++
			}
		}
		if preCurrentDevices == -1 {
			preCurrentDevices = currentDevices
		} else {
			if preCurrentDevices > 0 && currentDevices > 0 {
				sum += preCurrentDevices * currentDevices
				preCurrentDevices = currentDevices
			} else if preCurrentDevices == 0 && currentDevices > 0 {
				preCurrentDevices = currentDevices
			}
		}
	}
	return sum
}

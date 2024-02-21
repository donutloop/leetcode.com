package kata

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var employeeCount int
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			employeeCount++
		}
	}
	return employeeCount
}

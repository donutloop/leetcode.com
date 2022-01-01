package Day_of_the_Year

import "time"

func dayOfYear(date string) int {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}

	currentYear := d.Year()
	sum := d.Day()
	for {
		d = d.Add(-24 * time.Hour * time.Duration(d.Day()))
		if d.Year() != currentYear {
			break
		}
		sum += d.Day()
	}
	return sum
}

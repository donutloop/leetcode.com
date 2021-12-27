package kata

import "time"

const layoutISO = "2006-01-02"

func daysBetweenDates(date1 string, date2 string) int {
	dA, err := time.Parse(layoutISO, date2)
	if err != nil {
		panic(err)
	}

	dB, err := time.Parse(layoutISO, date1)
	if err != nil {
		panic(err)
	}
	d := int(dA.Sub(dB).Hours() / 24)
	if d < 0 {
		return d * -1
	}

	return d
}

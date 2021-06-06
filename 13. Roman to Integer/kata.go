package _3__Roman_to_Integer

var mapping = map[string]int{"I":1, "V": 5, "X":10, "L":50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}

func RomanToInt(s string) int {
	var sum int
	for i := 0; i < len(s); {

		if i+2 <= len(s) {
			v, ok := mapping[string(s[i:i+2])]
			if ok {
				sum += v
				i = i + 2
				continue
			}
		}
		sum += mapping[string(s[i])]
		i++

	}
	return sum
}

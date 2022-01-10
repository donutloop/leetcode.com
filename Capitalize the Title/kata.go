package Capitalize_the_Title

func capitalizeTitle(title string) string {

	b := make([]byte, len(title))
	var c int
	var j int
	for i := 0; i < len(title); i++ {

		if 'A' <= title[i] && title[i] <= 'Z' {
			b[i] = title[i] + 32
		} else {
			b[i] = title[i]
		}

		if title[i] == ' ' || len(title)-1 == i {
			if len(title)-1 == i {
				c++
			}

			if c > 2 && 'a' <= title[j] && title[j] <= 'z' {
				b[j] = title[j] - 32
			} else if c > 2 {
				b[j] = title[j]
			}
			c = 0
			j = i + 1
		} else {
			c++
		}
	}

	return string(b)
}

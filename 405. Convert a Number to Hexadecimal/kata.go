package kata

import "math/bits"

func ToHex(num int) string {
	if num == 0 {
		return "0"
	}

	n := uint(num)

	hex := make([]byte, 0)


	for {
		var r uint
		n, r = bits.Div( 0, n, 16)
		if r >= 0 && r <= 9 {
			hex = append(hex, 48+byte(r))
		} else if r >= 10 && r <= 15 {
			hex = append(hex, 97+byte(r)-10)
		}

		if len(hex) == 8 {
			break
		}

		if n == 0 {
			break
		}
	}

	j := len(hex)-1
	for i := 0; i < len(hex)/2; i++ {
		hex[i], hex[j] = hex[j], hex[i]
		j--
	}

	return string(hex)
}

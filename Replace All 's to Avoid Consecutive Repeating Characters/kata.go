package kata

import (
	"math/rand"
	"time"
)

func modifyString(s string) string {

	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if b[i] == '?' {
			if i-1 >= 0 {
				b[i] = pickChar(b[i-1])
				if i+1 < len(b) && b[i+1] != '?' && b[i] == b[i+1] {
					for {
						rand.Seed(time.Now().UnixNano())
						n := rand.Intn(26)
						if n == 0 {
							n++
						}
						b[i] = byte(n + 96)
						if b[i] != b[i-1] && b[i+1] != b[i] {
							break
						}
					}
				}
			} else if i+1 < len(b) {
				if b[i] == '?' {
					char := b[i+1]
					if b[i+1] == '?' {
						char = 97
					}

					b[i] = pickChar(char)
				}
			} else {
				b[i] = 97
			}
		}
	}

	return string(b)
}

func pickChar(char byte) byte {
	if char == 122 {
		return char - 1
	}
	if char == 97 {
		return char + 1
	}
	return char - 1
}

package _309__Decrypt_String_from_Alphabet_to_Integer_Mapping

func freqAlphabets(s string) string {

	ts := make([]byte, 0, len(s)/3*2)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			i--
			k := i
			i--
			j := i
			ts = append(ts, hashDecoder[s[j]][s[k]])
		} else {
			ts = append(ts, numberDecoder[s[i]])
		}
	}

	k := 0
	for i := len(ts) - 1; i >= (len(ts) / 2); i-- {
		ts[i], ts[k] = ts[k], ts[i]
		k++
	}

	return string(ts)
}

var numberDecoder map[byte]byte = map[byte]byte{'1': 'a', '2': 'b', '3': 'c', '4': 'd', '5': 'e', '6': 'f', '7': 'g', '8': 'h', '9': 'i'}

var hashDecoder map[byte]map[byte]byte = map[byte]map[byte]byte{
	'1': {'0': 'j', '1': 'k', '2': 'l', '3': 'm', '4': 'n', '5': 'o', '6': 'p', '7': 'q', '8': 'r', '9': 's'},
	'2': {'0': 't', '1': 'u', '2': 'v', '3': 'w', '4': 'x', '5': 'y', '6': 'z'},
}

package Decode_XORed_Array]

func decode(encoded []int, first int) []int {
	decoded := make([]int, len(encoded)+1)
	decoded[0] = first
	for i, n := range encoded {
		i = i+1
		inverse := first ^ n
		decoded[i] = inverse
		first = inverse
	}
	return decoded
}

package kata

func numJewelsInStones(J string, S string) int {
	set := make(map[byte]bool, len(J))
	for i := 0; i < len(J); i++ {
		set[J[i]] = true
	}
	var counter int
	for i := 0; i < len(S); i++ {
		ok := set[S[i]]
		if ok {
			counter++
		}
	}
	return counter
}

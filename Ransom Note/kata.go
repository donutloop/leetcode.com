package kata

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	set := make(map[byte]int)
	var i int
	for i = 0; i < len(magazine); i++ {
		set[magazine[i]]++
	}
	var c int
	var ok bool
	for i = 0; i < len(ransomNote); i++ {
		c, ok = set[ransomNote[i]]
		if !ok || c <= 0 {
			return false
		}
		set[ransomNote[i]]--
	}
	return true
}

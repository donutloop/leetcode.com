package kata

func reverseWords(s string) string {
	var word string
	var ss string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && word == "" {
			continue
		} else if s[i] == ' ' {
			ss = ss + " " + word
			if (i-1) > 0 && s[i-1] != ' ' {
				word = string(s[i-1])
				i--
			} else {
				word = ""
			}
			continue
		}
		word = string(s[i]) + word
	}
	if word != "" {
		ss = ss + " " + word
	}
	if ss == "" {
		return ss
	}
	return ss[1:]
}

package kata

type WordFilter struct {
	words []string
}

func Constructor(words []string) WordFilter {
	return WordFilter{words: words}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if prefix == "" && suffix == "" && len(this.words) != 0 {
		return 0
	}
	for i := len(this.words) - 1; i >= 0; i-- {
		if len(suffix) > len(this.words[i]) {
			continue
		}
		if len(prefix) > len(this.words[i]) {
			continue
		}

		if suffix != "" && prefix != "" {
			s := this.words[i][len(this.words[i])-len(suffix):]
			p := this.words[i][:len(prefix)]
			if s == suffix && p == prefix {
				return i
			}
		} else if suffix != "" {
			s := this.words[i][len(this.words[i])-len(suffix):]
			if s == suffix {
				return i
			}
			continue
		} else if prefix != "" {
			p := this.words[i][:len(prefix)]
			if p == prefix {
				return i
			}
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

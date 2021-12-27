package kata

type node struct {
	Val   rune
	word  bool
	nodes map[rune]*node
}

type Trie struct {
	node map[rune]*node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var current *node
	v, ok := this.node[rune(word[0])]
	if ok {
		current = v
	} else {
		if this.node == nil {
			this.node = make(map[rune]*node)
		}
		n := &node{Val: rune(word[0]), nodes: make(map[rune]*node)}
		this.node[rune(word[0])] = n
		current = n
	}
	for _, c := range word[1:] {
		v, ok := current.nodes[c]
		if ok {
			current = v
		} else {
			n := &node{Val: c, nodes: make(map[rune]*node)}
			current.nodes[c] = n
			current = n
		}
	}
	current.word = true
}

/* Returs if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this.node == nil {
		return false
	}
	if len(word) == 0 {
		return false
	}
	var current *node
	v, ok := this.node[rune(word[0])]
	if ok {
		current = v
	} else {
		return false
	}
	for _, c := range word[1:] {
		v, ok := current.nodes[c]
		if ok {
			current = v
		} else {
			return false
		}
	}
	return current.word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this.node == nil {
		return false
	}
	if len(prefix) == 0 {
		return false
	}
	var current *node
	v, ok := this.node[rune(prefix[0])]
	if ok {
		current = v
	} else {
		return false
	}
	for _, c := range prefix[1:] {
		v, ok := current.nodes[c]
		if ok {
			current = v
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

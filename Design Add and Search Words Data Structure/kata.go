package Design_Add_and_Search_Words_Data_Structure

type node struct {
	word  bool
	nodes [26]*node
}

type WordDictionary struct {
	rootNode *node
}

func Constructor() WordDictionary {
	return WordDictionary{
		rootNode: &node{nodes: [26]*node{}},
	}
}

func (this *WordDictionary) AddWord(word string) {
	currentNode := this.rootNode
	for i, char := range word {
		n := currentNode.nodes[int(char)-97]
		if n == nil {
			currentNode.nodes[int(char)-97] = &node{nodes: [26]*node{}, word: (len(word) - 1) == i}
			n = currentNode.nodes[int(char)-97]
		}
		currentNode = n
	}
}

func (this *WordDictionary) Search(word string) bool {
	return searchInNode(this.rootNode, word)
}

func searchInNode(rootNode *node, word string) bool {
	currentNode := rootNode
	for i, char := range word {
		if char == '.' {
			for _, n := range currentNode.nodes {
				if n == nil {
					continue
				}

				if (len(word)-1) == i && n.word {
					return true
				}

				ok := searchInNode(n, word[i+1:])
				if ok {
					return true
				}
			}
			return false
		} else {
			n := currentNode.nodes[int(char)-97]
			if n == nil {
				break
			}
			if (len(word)-1) == i && n.word {
				return true
			}
			currentNode = n
		}
	}
	return false
}

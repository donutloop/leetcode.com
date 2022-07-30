package kata

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {

	currentNode := root
	for currentNode != nil {
		if currentNode.Child != nil {
			r(currentNode, currentNode.Next, currentNode.Child)
			continue
		}
		currentNode = currentNode.Next
	}
	return root
}

func r(currentNode *Node, nextNode *Node, childNode *Node) {
	currentNode.Child = nil

	currentChildNode := childNode
	prevChildNode := childNode
	for currentChildNode != nil {
		if currentChildNode.Child != nil {
			r(currentChildNode, currentChildNode.Next, currentChildNode.Child)
			continue
		}
		prevChildNode = currentChildNode
		currentChildNode = currentChildNode.Next
	}

	currentNode.Next = childNode
	childNode.Prev = currentNode

	if nextNode != nil {
		nextNode.Prev = prevChildNode
		prevChildNode.Next = nextNode
	}
}

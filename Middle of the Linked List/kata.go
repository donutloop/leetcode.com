package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var current *ListNode
	nodes := make([]*ListNode, 0)
	current = head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}
	if len(nodes) == 1 {
		return nodes[0]
	}
	if len(nodes)%2 == 1 {
		return nodes[len(nodes)/2]
	}

	mid := (len(nodes) + 1) / 2
	if mid >= len(nodes) {
		mid = len(nodes) - 1
	}
	return nodes[mid]
}

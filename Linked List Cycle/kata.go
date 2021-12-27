package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	current := head.Next
	if current == nil {
		return false
	}
	count := map[*ListNode]struct{}{}
	for current != nil {
		_, ok := count[current]
		if ok {
			return true
		}
		count[current] = struct{}{}
		current = current.Next
	}
	return false
}

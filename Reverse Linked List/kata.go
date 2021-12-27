package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	current := head.Next
	next := &ListNode{Val: head.Val}
	for current != nil {
		head = &ListNode{Val: current.Val, Next: next}
		next = head
		current = current.Next
	}

	return head
}

package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && head.Val == val {
		return nil
	}
	current := head.Next
	pre := head
	for current != nil {
		if current.Val == val {
			pre.Next = current.Next
			current = current.Next
		} else {
			pre = current
			current = current.Next
		}
	}
	if head != nil && head.Val == val {
		return head.Next
	}
	return head
}

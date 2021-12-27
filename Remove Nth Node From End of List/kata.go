package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	current := head
	var max int
	for current != nil {
		max = max + 1
		current = current.Next
	}
	nth := max - n
	if nth == 0 {
		return head.Next
	}
	current = head.Next
	pre := head
	counter := 1
	for current != nil {
		if counter == nth {
			pre.Next = current.Next
			break
		}
		pre = current
		current = current.Next
		counter = counter + 1
	}
	return head
}

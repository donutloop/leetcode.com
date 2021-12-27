package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head.Next
	pre := head
	for current != nil && pre != nil {
		pre.Val, current.Val = current.Val, pre.Val
		pre = current.Next
		if current.Next != nil {
			current = current.Next.Next
		}
	}
	return head
}

package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := make(map[int]int)
	current := head
	for current != nil {
		count[current.Val] = count[current.Val] + 1
		current = current.Next
	}
	tmp := new(ListNode)
	current = head
	pre := tmp
	pre.Next = head
	for current != nil {
		v := count[current.Val]
		if v > 1 {
			pre.Next = current.Next
			current = current.Next
			continue
		}
		pre = current
		current = current.Next
	}
	return tmp.Next
}

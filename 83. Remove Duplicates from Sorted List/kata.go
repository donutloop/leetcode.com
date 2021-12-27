package _3__Remove_Duplicates_from_Sorted_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head.Next
	n := head
	for current != nil {
		if current.Val == n.Val {
			if current.Next != nil {
				n.Next = current.Next
				current = current.Next
			} else {
				n.Next = nil
				current = nil
			}
			continue
		}
		n = current
		current = current.Next
	}
	return head
}

package _1__Merge_Two_Sorted_Lists

type ListNode struct {
	Val int
	Next *ListNode
 }

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	currentA := l1
	currentB := l2
	root := &ListNode{}
	current := root
	for currentB != nil || currentA != nil {
		if currentB != nil && currentA != nil {
			if  currentB.Val <= currentA.Val {
				current.Next = currentB
				current = currentB
				currentB = currentB.Next
			} else if currentB.Val >= currentA.Val {
				current.Next = currentA
				current = currentA
				currentA = currentA.Next
			}
		} else if currentB != nil {
			current.Next = currentB
			currentB = nil
		} else if currentA != nil {
			current.Next = currentA
			currentA = nil
		}
	}
	return root.Next
}

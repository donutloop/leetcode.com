package _60__Intersection_of_Two_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	table := make(map[*ListNode]struct{}, 0)
	current := headA
	for current != nil {
		table[current] = struct{}{}
		current = current.Next
	}

	current = headB
	for current != nil {
		if _, ok := table[current]; ok {
			return current
		}
		current = current.Next
	}
	return nil
}

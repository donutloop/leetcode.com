package kata

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type NodeBacklink struct {
	*ListNode
	Prev *NodeBacklink
}

func getDecimalValue(head *ListNode) int {

	current := &NodeBacklink{ListNode: head}
	for current.Next != nil {
		current = &NodeBacklink{ListNode: current.Next, Prev: current}
	}

	var sum int
	if current != nil {
		if current.Val == 1 {
			sum += 1
		}
		current = current.Prev
	} else {
		return sum
	}

	step := 1
	for current != nil {
		step = step * 2
		if current.Val == 1 {
			sum += step
		}
		current = current.Prev
	}

	return sum
}

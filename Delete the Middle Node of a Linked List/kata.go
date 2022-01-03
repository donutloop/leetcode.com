package Delete_the_Middle_Node_of_a_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {

	if head.Next == nil {
		return nil
	}

	c := 0
	currentNode := head
	for currentNode != nil {
		c++
		currentNode = currentNode.Next
	}

	var previousNode *ListNode
	currentNode = head
	for i := 0; i < c/2; i++ {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	if currentNode.Next == nil {
		previousNode.Next = nil
	} else {
		*currentNode = *currentNode.Next
	}

	return head
}

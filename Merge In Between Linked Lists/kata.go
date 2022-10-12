package Merge_In_Between_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	rootNodeB := list2
	var tailNodeB *ListNode
	currentNodeB := list2
	for currentNodeB != nil {
		if currentNodeB.Next == nil {
			tailNodeB = currentNodeB
		}
		currentNodeB = currentNodeB.Next
	}

	var aNode, bNode *ListNode
	currentNode := list1
	var previousNode *ListNode
	var c int
	for currentNode != nil {
		if aNode == nil && c == a {
			aNode = previousNode
		}
		if bNode == nil && c == b {
			bNode = currentNode.Next
			break
		}
		previousNode = currentNode
		currentNode = currentNode.Next
		c++
	}

	aNode.Next = rootNodeB
	tailNodeB.Next = bNode

	return list1
}

package Merge_Nodes_in_Between_Zeros

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {

	currentNode := head

	var newChain *ListNode
	var currentNewChain *ListNode

	var sumNode *ListNode
	for currentNode != nil {

		if currentNode.Val == 0 {
			newSumNode := currentNode.Next
			if newSumNode == nil {
				sumNode.Next = nil
				break
			}
			sumNode = newSumNode

			if currentNode.Next != nil {
				currentNode = currentNode.Next.Next
			} else {
				break
			}

			if newChain == nil {
				newChain = sumNode
				currentNewChain = newChain
			} else {
				currentNewChain.Next = sumNode
				currentNewChain = currentNewChain.Next
			}

			continue
		} else {
			sumNode.Val += currentNode.Val
			sumNode.Next = currentNode.Next
			currentNode = currentNode.Next
		}

	}

	return newChain

}

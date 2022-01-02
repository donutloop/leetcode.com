package Find_the_Minimum_and_Maximum_Number_of_Nodes_Between_Critical_Points

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {

	if head == nil {
		panic("node list head is nil")
	}

	previousNode := head
	currentNode := head.Next

	c := 1
	matched := make([]int, 0)
	min := math.Inf(1)
	for currentNode != nil {
		if currentNode.Val < previousNode.Val && (currentNode.Next != nil && currentNode.Next.Val > currentNode.Val) {
			matched = append(matched, c)
		} else if currentNode.Val > previousNode.Val && (currentNode.Next != nil && currentNode.Next.Val < currentNode.Val) {
			matched = append(matched, c)
		}

		if len(matched) >= 2 && float64(matched[len(matched)-1]-matched[len(matched)-2]) < min {
			min = float64(matched[len(matched)-1] - matched[len(matched)-2])
		}

		previousNode = currentNode
		currentNode = currentNode.Next
		c++
	}

	if len(matched) == 0 || len(matched) == 1 {
		return []int{-1, -1}
	}

	return []int{int(min), matched[len(matched)-1] - matched[0]}
}

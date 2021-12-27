package _47__Insertion_Sort_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	nodes := make([]*ListNode, 0)
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}
	Insertionsort(nodes)
	return head
}

func Insertionsort(a []*ListNode) {
	for j := 1; j < len(a); j++ {
		key := a[j].Val
		i := j - 1
		for i >= 0 && a[i].Val > key {
			a[i+1].Val = a[i].Val
			i = i - 1
		}
		a[i+1].Val = key
	}
}

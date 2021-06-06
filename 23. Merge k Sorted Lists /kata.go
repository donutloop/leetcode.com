package _3__Merge_k_Sorted_Lists_

import "container/heap"

type ListNode struct {
	   Val int
	     Next *ListNode
}

type Item struct {
	value    *ListNode
	priority int
	index int
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	pq := make(PriorityQueue, 0)
	i := 0
	var current *ListNode
	for _, root := range lists {
		current = root
		for current != nil {
			item := &Item{
				value:    current,
				priority: current.Val,
				index:    i,
			}
			pq = append(pq, item)
			i++
			current = current.Next
		}
	}
	if pq.Len() == 0 {
		return nil
	}
	heap.Init(&pq)
	root := heap.Pop(&pq).(*Item).value
	current = root
	for pq.Len() > 0 {
		current.Next =  heap.Pop(&pq).(*Item).value
		current = current.Next
	}
	current.Next = nil
	return root
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

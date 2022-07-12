package Remove_Stones_to_Minimize_the_Total

import (
	"container/heap"
	"math"
	"sort"
)

func minStoneSum(piles []int, k int) int {

	h := IntHeap(piles)
	sort.Sort(h)

	for i := 0; i < k; i++ {
		n := heap.Pop(&h).(int)
		n = n - int(math.Floor(float64(n)/2))
		heap.Push(&h, n)
	}

	var sum int
	for i := 0; i < h.Len(); i++ {
		sum += h[i]
	}

	return sum
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {

	h[j], h[i] = h[i], h[j]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	if n == 0 {
		return -1
	}
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

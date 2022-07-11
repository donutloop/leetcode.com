package Largest_Number_After_Digit_Swaps_by_Parity

import (
	"container/heap"
	"math"
)

func largestInteger(num int) int {

	odd := &IntHeap{}
	even := &IntHeap{}

	var i int

	var bitSet int

	for {
		n := num % 10
		num = num / 10
		if n%2 == 0 {
			even.Push(n)
			if bitSet&(1<<i) == 0 {
				bitSet |= 1 << i
			}
		} else {
			odd.Push(n)
			if bitSet&(1<<i) > 0 {
				bitSet ^= 1 << i
			}
		}

		if num == 0 {
			break
		}

		i++
	}
	heap.Init(odd)
	heap.Init(even)

	num = 0
	u := int(math.Pow(10, float64(i)))
	for ; i >= 0; i-- {

		if (bitSet & (1 << i)) == 0 {
			num += heap.Pop(odd).(int) * u
		} else {
			num += heap.Pop(even).(int) * u
		}

		u = u / 10
	}

	return num
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	if h.Len() < 2 {
		return
	}

	h[i], h[j] = h[j], h[i]
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

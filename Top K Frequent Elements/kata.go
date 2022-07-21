package kata

import "container/heap"

type Item struct {
	value    int
	priority int
	index    int
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	pq := make(PriorityQueue, 0)
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n] = counter[n] + 1
	}
	var i int
	for n, c := range counter {
		item := &Item{
			value:    n,
			priority: c,
			index:    i,
		}
		pq = append(pq, item)
		i++
	}
	nums = nil
	heap.Init(&pq)
	for pq.Len() > 0 && k > 0 {
		nums = append(nums, heap.Pop(&pq).(*Item).value)
		k--
	}
	return nums
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority > pq[j].priority {
		return true
	}
	return false
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


SELECT ((SUM(CASE WHEN activity_type = "open_session" THEN 1 ELSE 0 END)-(SUM(CASE WHEN activity_type = "open_session" THEN -1 ELSE 0 END)-SUM(CASE WHEN activity_type = "end_session" THEN -1 ELSE 0 END)))) AS a
FROM Activity
WHERE (activity_type = "open_session" OR activity_type = "end_session")
AND activity_date BETWEEN DATE_SUB("2019-07-27", INTERVAL 30 DAY) AND "2019-07-27" GROUP BY user_id;
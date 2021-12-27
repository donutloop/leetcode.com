package kata

import "container/list"

func processQueries(queries []int, m int) []int {

	l := list.New()

	for i := 1; i < m+1; i++ {
		l.PushBack(i)
	}

	idx := make([]int, len(queries))
	j := 0
	for _, q := range queries {
		var i int
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == q {
				l.MoveToFront(e)
				idx[j] = i
				j++
				break
			}
			i++
		}
	}

	return idx
}

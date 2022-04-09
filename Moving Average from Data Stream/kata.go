package Moving_Average_from_Data_Stream

import "container/list"

type MovingAverage struct {
	currentLen int
	size       int
	list       *list.List
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		list: list.New(),
	}
}

func (mv *MovingAverage) Next(val int) float64 {
	mv.currentLen++
	mv.list.PushBack(val)
	if mv.currentLen > mv.size {
		mv.list.Remove(mv.list.Front())
		mv.currentLen--
	}
	var avg float64
	for e := mv.list.Front(); e != nil; e = e.Next() {
		avg += float64(e.Value.(int))
	}
	return avg / float64(mv.currentLen)
}

package Moving_Average_from_Data_Stream

import "container/list"

type MovingAverage struct {
	currentLen int
	size       int
	list       *list.List
	sum        float64
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
		n := mv.list.Front().Value.(int)
		mv.list.Remove(mv.list.Front())
		mv.sum = mv.sum - float64(n)
		mv.currentLen--
	}
	mv.sum = mv.sum + float64(val)
	return mv.sum / float64(mv.currentLen)
}

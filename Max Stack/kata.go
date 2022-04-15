package Max_Stack

import (
	"container/list"
	"math"
)

type MaxStack struct {
	list *list.List
	max  *list.Element
}

func Constructor() MaxStack {
	return MaxStack{list: list.New()}
}

func (this *MaxStack) Push(x int) {
	e := this.list.PushBack(x)
	if this.max == nil {
		this.max = e
	} else if x >= this.max.Value.(int) {
		this.max = e
	}
}

func (this *MaxStack) Pop() int {
	e := this.list.Back()
	if e == nil {
		return -1
	}
	this.list.Remove(e)
	if e == this.max {
		this.max = nil
		max := math.MaxInt64
		for e1 := this.list.Front(); e1 != nil; e1 = e1.Next() {
			if max == math.MaxInt64 || max <= e1.Value.(int) {
				max = e1.Value.(int)
				this.max = e1
			}
		}
	}
	return e.Value.(int)
}

func (this *MaxStack) Top() int {
	e := this.list.Back()
	if e == nil {
		return -1
	}
	return e.Value.(int)
}

func (this *MaxStack) PeekMax() int {
	return this.max.Value.(int)
}

func (this *MaxStack) PopMax() int {
	e := this.max
	this.list.Remove(e)
	max := math.MaxInt64
	this.max = nil
	for e1 := this.list.Front(); e1 != nil; e1 = e1.Next() {
		if max == math.MaxInt64 || max <= e1.Value.(int) {
			max = e1.Value.(int)
			this.max = e1
		}
	}
	return e.Value.(int)
}

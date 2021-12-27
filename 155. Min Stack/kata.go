package _55__Min_Stack

type MinStack struct {
	list []int
}

func Constructor() MinStack {
	return MinStack{
		list: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
}

func (this *MinStack) Pop() {
	if len(this.list) == 0 {
		return
	}
	this.list = this.list[:len(this.list)-1]
}

func (this *MinStack) Top() int {
	if len(this.list) == 0 {
		return -1
	}
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	var min *int
	for i := 0; i < len(this.list); i++ {
		if min == nil || *min > this.list[i] {
			if min == nil {
				min = new(int)
			}
			*min = this.list[i]
		}
	}
	return *min
}

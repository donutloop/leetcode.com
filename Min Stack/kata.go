package kata

type MinStack struct {
	list []int
}

/** initialize your data structure here. */
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

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

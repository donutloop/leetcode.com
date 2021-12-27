package kata

type FreqStack struct {
	counter map[int]int
	list    []int
}

func Constructor() FreqStack {
	return FreqStack{
		counter: make(map[int]int),
		list:    make([]int, 0),
	}
}

func (this *FreqStack) Push(x int) {
	this.counter[x] = this.counter[x] + 1
	this.list = append(this.list, x)
}

func (this *FreqStack) Pop() int {
	if len(this.list) == 1 {
		v := this.list[0]
		this.list = this.list[1:]
		return v
	}
	max := 0
	index := 0
	for i := len(this.list) - 1; i > 0; i-- {
		c := this.counter[this.list[i]]
		if c > max {
			max = c
			index = i
		}
	}
	v := this.list[index]
	this.counter[v] = this.counter[v] - 1
	this.list = append(this.list[:index], this.list[index+1:]...)
	return v
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */

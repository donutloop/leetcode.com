package kata

type CustomStack struct {
	elements []int
	maxSize  int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		elements: make([]int, 0),
		maxSize:  maxSize,
	}
}

func (this *CustomStack) Push(x int) {

	if len(this.elements) >= this.maxSize {
		return
	}

	this.elements = append(this.elements, x)
}

func (this *CustomStack) Pop() int {
	if len(this.elements) == 0 {
		return -1
	}

	x := this.elements[len(this.elements)-1]
	this.elements = this.elements[:len(this.elements)-1]

	return x
}

func (this *CustomStack) Increment(k int, val int) {

	if k >= len(this.elements) {
		k = len(this.elements)
	}

	for i := 0; i < k; i++ {
		this.elements[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

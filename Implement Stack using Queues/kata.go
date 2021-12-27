package kata

type MyStack struct {
	values []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		values: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.values = append(this.values, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.values[len(this.values)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.values) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

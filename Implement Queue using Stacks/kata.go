package kata

type MyQueue struct {
	list []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		list: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.list = append([]int{x}, this.list...)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	v := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.list[len(this.list)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.list) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

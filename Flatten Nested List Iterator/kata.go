package kata

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	c            chan int
	currentValue *int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	c := make(chan int)
	go recursiveChannling(nestedList, c, true)
	return &NestedIterator{c: c}
}

func (this *NestedIterator) Next() int {
	if this.currentValue == nil {
		panic("unknown order of operation")
	}
	v := this.currentValue
	this.currentValue = nil
	return *v
}

func (this *NestedIterator) HasNext() bool {
	v, ok := <-this.c
	if ok {
		pv := new(int)
		*pv = v
		this.currentValue = pv
		return true
	}
	return false
}

func recursiveChannling(nestedList []*NestedInteger, c chan int, firstLayer bool) {
	for _, o := range nestedList {
		if o.IsInteger() {
			c <- o.GetInteger()
		} else {
			recursiveChannling(o.GetList(), c, false)
		}
	}

	if firstLayer {
		close(c)
	}
}

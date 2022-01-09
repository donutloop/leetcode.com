package Design_Linked_List

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	rootNode *Node
	lenght   int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.lenght {
		return -1
	}

	var i int
	currentNode := this.rootNode
	for currentNode != nil {
		if index == i {
			return currentNode.val
		}
		currentNode = currentNode.next
		i++
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.lenght++
	nextNode := this.rootNode
	this.rootNode = &Node{
		val:  val,
		next: nextNode,
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.rootNode == nil {
		this.AddAtHead(val)
		return
	}

	currentNode := this.rootNode
	for currentNode != nil {
		if currentNode.next != nil {
			currentNode = currentNode.next
			continue
		} else {
			break
		}
	}
	currentNode.next = &Node{val: val}
	this.lenght++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if this.rootNode == nil {
		return
	}

	currentNode := this.rootNode
	var i int
	for currentNode != nil {
		if index-1 == i {
			break
		} else if currentNode.next != nil {
			currentNode = currentNode.next
			i++
			continue
		} else {
			break
		}
	}

	nextNode := currentNode.next
	currentNode.next = &Node{
		val:  val,
		next: nextNode,
	}
	this.lenght++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.lenght {
		return
	}

	if this.lenght == 0 {
		return
	}

	if index == 0 {
		this.rootNode = this.rootNode.next
		this.lenght--
		return
	}

	currentNode := this.rootNode
	var i int
	for currentNode != nil {
		if currentNode.next != nil {
			previousNode := currentNode
			i++
			if i == index {
				previousNode.next = currentNode.next.next
				break
			}
			currentNode = currentNode.next
			continue
		} else {
			break
		}
	}

	this.lenght--

	return
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
